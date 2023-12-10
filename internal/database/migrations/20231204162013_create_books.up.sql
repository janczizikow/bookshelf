CREATE TABLE
  IF NOT EXISTS books (
    id serial PRIMARY KEY,
    title text NOT NULL DEFAULT '',
    ISBN text NOT NULL,
    price numeric(14,2) NOT NULL,
    publisher text NOT NULL,
    pages integer NOT NULL,
    publish_date date NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT now (),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT now (),
    author_id integer NOT NULL,
    CONSTRAINT books_author_id_fk FOREIGN KEY (author_id) REFERENCES authors (id)
  );

CREATE INDEX IF NOT EXISTS idx_books_title ON books (title);

CREATE UNIQUE INDEX IF NOT EXISTS idx_books_isbn ON books (ISBN)
WHERE ISBN <> '';