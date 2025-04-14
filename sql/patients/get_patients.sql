SELECT
    id,
    first_name,
    last_name,
    middle_name,
    birth_date,
    phone_number,
    email,
    version
FROM patients
WHERE deleted_ts ISNULL
ORDER BY last_name
LIMIT $1
OFFSET $2;