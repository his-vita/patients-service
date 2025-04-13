-- Active: 1744043592357@@127.0.0.1@5432@patients_db
-- +goose Up
-- +goose StatementBegin
CREATE TABLE patients (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100),
    birth_date DATE NOT NULL,
    phone_number VARCHAR(11),
    email VARCHAR(64),
    created_ts TIMESTAMP DEFAULT NULL,
    created_by VARCHAR(64) DEFAULT NULL,
    updated_ts TIMESTAMP DEFAULT NULL,
    updated_by VARCHAR(64) DEFAULT NULL,
    deleted_ts TIMESTAMP DEFAULT NULL,
    deleted_by VARCHAR(64) DEFAULT NULL,
    version int DEFAULT 0
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE patients;
-- +goose StatementEnd