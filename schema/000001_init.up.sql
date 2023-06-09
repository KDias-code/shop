CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       number VARCHAR(255) UNIQUE NOT NULL,
                       name VARCHAR(255) NOT NULL,
                       surname VARCHAR(255) NOT NULL,
                       password VARCHAR(255) NOT NULL
);