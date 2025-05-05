UPDATE contacts
SET
    phone_number = $2,
    work_phone_number = $3,
    email = $4
WHERE
    patient_id = $1