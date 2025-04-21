UPDATE patients
SET
    first_name = $2,
    last_name = $3,
    middle_name = $4,
    birth_date = $5,
    gender = $6,
    phone_number = $7,
    email = $8,
    updated_ts = NOW(),
    updated_by = $9,
    version = version + 1
WHERE
    id = $1 AND
    version = $10;