INSERT INTO
    patients (
        first_name,
        last_name,
        middle_name,
        birth_date,
        gender,
        created_ts,
        created_by
    )
VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, $6)
RETURNING id;