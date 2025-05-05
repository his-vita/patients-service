-- +goose Up
-- +goose StatementBegin
CREATE TABLE personal_documents(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    series CHAR(4) NOT NULL,
    number CHAR(6) NOT NULL,
    department_code CHAR(6) NOT NULL,
    issue_date DATE NOT NULL,
    expiry_date DATE DEFAULT NULL,
    main BOOLEAN NOT NULL,
    patient_id UUID NOT NULL,
    document_type_id INT,
    document_company_id INT,
    FOREIGN KEY (patient_id) REFERENCES patients(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
