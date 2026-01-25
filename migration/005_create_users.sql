CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    role_id INTEGER REFERENCES roles(id)
);