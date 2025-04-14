UPDATE patients
SET
    first_name = $2,
    last_name = $3,
    middle_name = $4,
    birth_date = $5,
    phone_number = $6,
    email = $7,
    updated_ts = NOW(),
    updated_by = $8,
    version = version + 1
WHERE
    id = $1 AND
    version = $9;