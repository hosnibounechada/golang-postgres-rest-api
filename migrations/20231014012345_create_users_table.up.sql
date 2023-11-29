CREATE TABLE IF NOT EXISTS users
(
    id serial PRIMARY KEY,
    username character varying(255) NOT NULL,
    first_name character varying(50) NOT NULL,
    last_name character varying(50) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255),
    verified boolean DEFAULT false,
    CONSTRAINT users_email_ukey UNIQUE (email),
    CONSTRAINT users_username_ukey UNIQUE (username)
);

CREATE TABLE IF NOT EXISTS devices
(
    id serial PRIMARY KEY,
    name character varying(255),
    os character varying(50),
    browser character varying(50),
    token character varying(128) NOT NULL,
    user_id integer,
    CONSTRAINT devices_token_ukey UNIQUE (token),
    CONSTRAINT devices_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
)