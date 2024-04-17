CREATE TABLE users (
  "id" uuid PRIMARY KEY,
  "mail" varchar UNIQUE,
  "first_name" varchar NOT NULL,
  "last_name" varchar,
  "password" varchar NOT NULL,
  "phone" varchar UNIQUE,
  "sex" varchar NOT NULL,
  "active" bool NOT NULL DEFAULT true,
  "created_at" TIMESTAMP DEFAULT NOW(),
  "updated_at" timestamp
);