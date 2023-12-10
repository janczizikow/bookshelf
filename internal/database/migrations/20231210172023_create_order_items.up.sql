CREATE TABLE
  IF NOT EXISTS order_items (
    id serial PRIMARY KEY,
    order_id integer NOT NULL,
    book_id integer NOT NULL,
    quantity integer NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT now (),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT now (),
    CONSTRAINT order_items_order_id_fk FOREIGN KEY (order_id) REFERENCES orders (id),
    CONSTRAINT order_items_book_id_fk FOREIGN KEY (book_id) REFERENCES books (id)
  );