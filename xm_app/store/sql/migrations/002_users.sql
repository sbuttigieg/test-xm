-- +migrate Up

-- users
CREATE TABLE IF NOT EXISTS users (
    id              UUID            NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name            VARCHAR(30)     NOT NULL UNIQUE,
    username        VARCHAR(30)     NOT NULL UNIQUE,
    email           VARCHAR(50)     NOT NULL UNIQUE,
    password        VARCHAR(100)    NOT NULL,
    created_at      TIMESTAMP       NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP       NOT NULL DEFAULT NOW()
);

INSERT INTO public.users (id, name, username, email, password) 
VALUES('28d09c66-d5a5-42ce-b70b-110a73a0c38f'::uuid, 'Joe Doe', 'divinity', 'thedoe@gmail.com', '$2a$14$P8ppRoUz6Q2wqMEDzuH8PePUYfejH1bZTkJsD4pWGNhjB9b86abRm');

-- -- +migrate Down
-- DROP TABLE IF EXISTS users;