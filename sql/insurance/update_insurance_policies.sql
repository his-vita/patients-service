UPDATE insurance_policies
SET
    number = $2,
    issue_date = $3,
    expiry_date = $4,
    type = $5,
    insurance_company_id = $6
WHERE
    id = $1