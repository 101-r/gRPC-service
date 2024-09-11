include .env
LOCAL_BIN:=$(PWD)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/user
	protoc --proto_path api/user \
	--go_out=pkg/user --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user/user.proto

make-test:
	test-create
	test-get
	test-delete

run-grpc:
	go run ./cmd/grpc_server/main.go

test-create:
	grpcurl -plaintext -d "{\"info\": {\"username\": \"root\", \"first_name\": \"Ravshan\", \"last_name\": \"Zaripov\", \"email\": \"test@example.com\", \"password\": \"secret\"}}" localhost:50051 user.UserService/Create

test-get:
	grpcurl -plaintext -d "{\"id\": 1}" localhost:50051 user.UserService/Get

test-delete:
	grpcurl -plaintext -d "{"\id\": 1}" localhost:50051 user.UserService/Delete