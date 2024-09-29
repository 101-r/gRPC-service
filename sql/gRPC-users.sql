CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "user_info_id" integer
);

CREATE TABLE "users_info" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "first_name" varchar,
  "last_name" varchar,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp DEFAULT now(),
  "updated_at" timestamp DEFAULT NULL
);

ALTER TABLE "users" ADD FOREIGN KEY ("id") REFERENCES "users_info" ("id");
