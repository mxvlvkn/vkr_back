CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(40) UNIQUE NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(40) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(40) NOT NULL,
    surname VARCHAR(40) NOT NULL,
    patronymic VARCHAR(40) NOT NULL,
    role_id INT NOT NULL REFERENCES roles(id)
);

CREATE TABLE units (
    id SERIAL PRIMARY KEY,
    code INTEGER UNIQUE NOT NULL,
    sign VARCHAR(10) NOT NULL,
    name VARCHAR(40) NOT NULL
);

CREATE UNIQUE INDEX idx_users_login ON users (LOWER(login));
CREATE INDEX idx_users_role ON users(role_id);