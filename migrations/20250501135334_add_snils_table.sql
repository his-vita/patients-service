-- +goose Up
-- +goose StatementBegin
CREATE TABLE snils (
    patient_id UUID PRIMARY KEY,
    number CHAR(11),
    FOREIGN KEY (patient_id) REFERENCES patients(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE snils;
-- +goose StatementEnd
