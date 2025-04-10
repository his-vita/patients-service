-- +goose Up
-- +goose StatementBegin
CREATE TABLE patients (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100),
    birth_date DATE NOT NULL,
    phone_number VARCHAR(11),
    email VARCHAR(64)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE patients;
-- +goose StatementEnd