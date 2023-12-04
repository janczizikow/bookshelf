CREATE TABLE
  IF NOT EXISTS authors (id serial PRIMARY KEY, "name" text NOT NULL);

CREATE UNIQUE INDEX IF NOT EXISTS idx_authors_name ON authors ("name")
WHERE
  "name" <> '';