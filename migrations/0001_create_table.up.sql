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

CREATE TABLE manufacturers (
    id SERIAL PRIMARY KEY,
    name        VARCHAR(60) NOT NULL,
    country     VARCHAR(40) NOT NULL,
    inn         VARCHAR(12) UNIQUE NOT NULL,
    ur_address   VARCHAR(200) NOT NULL,
    fact_address VARCHAR(200) NOT NULL,
    fio         VARCHAR(200) NOT NULL,
    phone       VARCHAR(20) NOT NULL,
    email       VARCHAR(200) NOT NULL
);

CREATE TABLE numenclatures (
    id SERIAL PRIMARY KEY,
    name        VARCHAR(400) NOT NULL,
    image_url   VARCHAR(200) NOT NULL,
    use_serial     BOOLEAN NOT NULL,
    use_marks     BOOLEAN NOT NULL,
    article         VARCHAR(200) UNIQUE NOT NULL,
    unit_id   INT NOT NULL REFERENCES units(id),
    manufacturer_id   INT NOT NULL REFERENCES manufacturers(id)
);

CREATE TABLE barcodes (
    id SERIAL PRIMARY KEY,
    code        VARCHAR(100) UNIQUE NOT NULL,
    numenclature_id   INT NOT NULL REFERENCES numenclatures(id)
);

CREATE TABLE marks (
    id SERIAL PRIMARY KEY,
    code        VARCHAR(200) UNIQUE NOT NULL,
    numenclature_id   INT NOT NULL REFERENCES numenclatures(id)
);

CREATE UNIQUE INDEX idx_users_login ON users (LOWER(login));
CREATE INDEX idx_users_role ON users(role_id);