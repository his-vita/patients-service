-- +goose Up
-- +goose StatementBegin
CREATE TABLE contacts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    patient_id UUID NOT NULL REFERENCES patients (id),
    phone_number VARCHAR(11) DEFAULT NULL,
    work_phone_number VARCHAR(11) DEFAULT NULL,
    email VARCHAR(32) DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
