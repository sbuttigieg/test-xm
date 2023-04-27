-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- companies
CREATE TABLE IF NOT EXISTS companies (
    id              UUID            NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name            VARCHAR(15)     NOT NULL UNIQUE,
    description     VARCHAR(3000),
    employees       INT             NOT NULL,
    registered      BOOLEAN         NOT NULL DEFAULT false,
    type            VARCHAR(30)     NOT NULL,
    created_at      TIMESTAMP       NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP       NOT NULL DEFAULT NOW()
);

INSERT INTO public.companies (id, name, description, employees, registered, type) 
VALUES('18d09c66-d5a5-42ce-b70b-110a73a0c38f'::uuid, 'Google', 'Home of Golang', 50000, true, 'Corporations');

-- -- +migrate Down
-- DROP TABLE IF EXISTS companies;