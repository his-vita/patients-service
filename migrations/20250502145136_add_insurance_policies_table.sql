-- +goose Up
-- +goose StatementBegin
CREATE TABLE insurance_policies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    number CHAR(16) NOT NULL,
    issue_date DATE NOT NULL,
    expiry_date DATE DEFAULT NULL,
    type INT,
    main BOOLEAN NOT NULL,
    patient_id UUID NOT NULL,
    insurance_company_id INT,
    FOREIGN KEY (patient_id) REFERENCES patients(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE insurance_policies;
-- +goose StatementEnd
