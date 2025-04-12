INSERT INTO patients (
    first_name,
    last_name,
    middle_name,
    birth_date,
    phone_number,
    email
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
);
