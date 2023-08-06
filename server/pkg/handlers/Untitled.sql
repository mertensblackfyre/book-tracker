CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" text,
  "role" text,
  "created_at" datetime
);

CREATE TABLE "books" (
  "id" integer PRIMARY KEY,
  "title" text,
  "author" text,
  "user_id" integer,
  "status" text,
  "price" real,
  "created_at" datetime
);

ALTER TABLE "books" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
