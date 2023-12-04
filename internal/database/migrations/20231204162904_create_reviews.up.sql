CREATE TABLE
  IF NOT EXISTS reviews (
    id serial PRIMARY KEY,
    rating smallint NOT NULL,
    "description" text NOT NULL DEFAULT '',
    book_id integer NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp(0)
    with
      time zone NOT NULL DEFAULT now (),
      updated_at timestamp(0)
    with
      time zone NOT NULL DEFAULT now (),
      CONSTRAINT reviews_book_id_fk FOREIGN KEY (book_id) REFERENCES books (id),
      CONSTRAINT reviews_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id)
  );