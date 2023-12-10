CREATE TABLE
  IF NOT EXISTS authors (
    id serial PRIMARY KEY, 
    "name" text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT now (),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT now ()
  );

CREATE UNIQUE INDEX IF NOT EXISTS idx_authors_name ON authors ("name")
WHERE "name" <> '';