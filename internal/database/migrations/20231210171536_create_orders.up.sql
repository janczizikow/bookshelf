CREATE TABLE
  IF NOT EXISTS orders (
    id serial PRIMARY KEY,
    user_id integer NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT now (),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT now (),
    total numeric(14,2) NOT NULL,
    CONSTRAINT orders_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id)
  );