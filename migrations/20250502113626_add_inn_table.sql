-- +goose Up
-- +goose StatementBegin
CREATE TABLE inn (
    patient_id UUID PRIMARY KEY,
    number CHAR(12),
    FOREIGN KEY (patient_id) REFERENCES patients(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE inn;
-- +goose StatementEnd
