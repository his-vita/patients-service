-- +goose Up
-- +goose StatementBegin
CREATE TABLE contacts (
    patient_id UUID PRIMARY KEY,
    phone_number CHAR(11) DEFAULT NULL,
    work_phone_number CHAR(11) DEFAULT NULL,
    email VARCHAR(32) DEFAULT NULL,
    FOREIGN KEY (patient_id) REFERENCES patients(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE contacts;
-- +goose StatementEnd
