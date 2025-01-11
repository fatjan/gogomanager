-- +goose Up
CREATE TABLE IF NOT EXISTS employees (
    id BIGSERIAL PRIMARY KEY,
    identity_number VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    employee_image_uri TEXT,
    gender VARCHAR(10),
    department_id BIGINT,
    manager_id BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (department_id) REFERENCES departments(id),
    FOREIGN KEY (manager_id) REFERENCES managers(id)
);
CREATE INDEX IF NOT EXISTS idx_employees_manager_id ON employees(manager_id);

-- COMPOSITE UNIQUE CONSTRAINT
ALTER TABLE employees 
ADD CONSTRAINT employees_identity_number_per_manager
UNIQUE (identity_number, manager_id);

-- +goose Down
DROP TABLE IF EXISTS employees;
DROP INDEX IF EXISTS idx_employees_manager_id;
