INSERT INTO
    patients (
        first_name,
        last_name,
        middle_name,
        birth_date,
        phone_number,
        email,
        created_ts,
        created_by
    )
VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, $7);