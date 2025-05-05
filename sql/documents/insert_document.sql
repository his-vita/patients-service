INSERT INTO
    personal_documents (
        patient_id,
        series,
        number,
        department_code,
        issue_date,
        expiry_date,
        main,
        document_type_id,
        document_company_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);