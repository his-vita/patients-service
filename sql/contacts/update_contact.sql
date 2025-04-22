UPDATE contacts
SET
    phone_number = $2,
    email = $3,
    main = $4
WHERE
    id = $1