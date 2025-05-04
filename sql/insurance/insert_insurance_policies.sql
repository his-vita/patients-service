INSERT INTO
    insurance_policies (
        patient_id,
        number,
        issue_date,
        expiry_date,
        type,
        insurance_company_id
    )
VALUES ($1, $2, $3, $4, $5, $6);