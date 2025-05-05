UPDATE personal_documents
SET
    series = $2,
    number = $3,
    department_code = $4,
    issue_date = $5,
    expiry_date = $6,
    main = $7,
    document_type_id = $8,
    document_company_id = $9
WHERE
    id = $1