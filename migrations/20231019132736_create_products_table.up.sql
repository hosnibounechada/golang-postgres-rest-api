CREATE TABLE IF NOT EXISTS products
(
    id serial PRIMARY KEY,
    name character varying(255) NOT NULL,
    price numeric NOT NULL,
    quantity numeric NOT NULL,
    user_id integer,
    CONSTRAINT products_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT products_price_check CHECK (price > 0::numeric),
    CONSTRAINT products_quantity_check CHECK (quantity > 0::numeric)
)