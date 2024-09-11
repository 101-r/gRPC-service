package user

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/101-r/gRPC-service/internal/model"
	"github.com/101-r/gRPC-service/internal/repository/converter"
	"github.com/101-r/gRPC-service/pkg/db/postgresql"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	def "github.com/101-r/gRPC-service/internal/repository"
	repoModel "github.com/101-r/gRPC-service/internal/repository/model"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	db *pgxpool.Pool
	m  sync.RWMutex
}

func NewRepository(connStr string) *repository {
	db, err := postgresql.New(connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &repository{
		db: db,
	}
}

func (r *repository) Create(_ context.Context, info *model.UserInfo) (int, error) {
	r.m.Lock()
	defer r.m.Unlock()

	tx, err := r.db.BeginTx(context.TODO(), pgx.TxOptions{})
	if err != nil {
		return -1, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if err != nil {
			errRb := tx.Rollback(context.TODO())
			if errRb != nil {
				err = fmt.Errorf("rollback failed: %w", errRb)
				return
			}
		}

		err = tx.Commit(context.TODO())
	}()

	return r.create(tx, info)
}

func (r *repository) create(tx pgx.Tx, info *model.UserInfo) (int, error) {
	var query string = `INSERT INTO users_info (username, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int

	err := tx.QueryRow(context.TODO(), query, info.Username, info.FirstName, info.LastName, info.Email, info.Password).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("fetching id: %w", err)
	}

	query = `INSERT INTO users (id) VALUES ($1)`

	res, err := tx.Exec(context.TODO(), query, id)
	if err != nil {
		return -1, fmt.Errorf("upsert: %w", err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return -1, fmt.Errorf("user id is not inserted")
	}

	return id, nil
}

func (r *repository) Get(_ context.Context, id int) (*model.UserInfo, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var query string = `SELECT * FROM users_info WHERE id = $1`
	var info repoModel.UserInfo

	err := r.db.QueryRow(context.TODO(), query, id).Scan(
		&info.Id,
		&info.Username,
		&info.FirstName,
		&info.LastName,
		&info.Email,
		&info.Password,
		&info.CreatedAt,
		&info.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("%w user is not found", err)
		}
		return nil, fmt.Errorf("%w removing user", err)
	}

	return converter.ToUserInfoFromRepo(&info), nil
}

func (r *repository) Delete(_ context.Context, id int) (int, error) {
	r.m.Lock()
	defer r.m.Unlock()

	tx, err := r.db.BeginTx(context.TODO(), pgx.TxOptions{})
	if err != nil {
		return -1, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if err != nil {
			errRb := tx.Rollback(context.TODO())
			if errRb != nil {
				err = fmt.Errorf("rollback failed: %w", errRb)
				return
			}
		}

		err = tx.Commit(context.TODO())
	}()

	return r.delete(tx, id)
}

func (r *repository) delete(tx pgx.Tx, id int) (int, error) {
	var query string = `DELETE FROM users WHERE id = $1`

	res, err := tx.Exec(context.TODO(), query, id)
	if err != nil {
		return -1, fmt.Errorf("failed to delete from users table: %w", err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return -1, fmt.Errorf("no user found in users table with id: %d", id)
	}

	query = `DELETE FROM users_info WHERE id = $1`

	res, err = tx.Exec(context.TODO(), query, id)
	if err != nil {
		return -1, fmt.Errorf("failed to delete from users_info table: %w", err)
	}

	rowsAffected = res.RowsAffected()
	if rowsAffected == 0 {
		return -1, fmt.Errorf("no user found in users_info table with id: %d", id)
	}

	return id, nil
}
