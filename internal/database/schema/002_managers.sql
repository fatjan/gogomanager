-- +goose Up
CREATE TABLE IF NOT EXISTS managers (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    user_image_uri TEXT,
    company_name VARCHAR(255),
    company_image_uri TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_managers_email ON managers(email);

-- +goose Down
DROP TABLE DROP TABLE IF EXISTS managers;
DROP INDEX IF EXISTS idx_managers_email;
