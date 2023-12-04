CREATE EXTENSION IF NOT EXISTS "citext";

CREATE TABLE
  IF NOT EXISTS users (
    id serial PRIMARY KEY,
    email citext NOT NULL,
    password_hash bytea NOT NULL,
    "name" text NOT NULL DEFAULT '',
    created_at timestamp(0)
    with
      time zone NOT NULL DEFAULT now (),
      updated_at timestamp(0)
    with
      time zone NOT NULL DEFAULT now (),
      deleted_at timestamp(0)
    with
      time zone
  );

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users (email)
WHERE
  email <> '';