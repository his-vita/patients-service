SELECT
    p.id,
    p.first_name,
    p.last_name,
    p.middle_name,
    p.birth_date,
    p.gender,
    c.phone_number,
    c.work_phone_number,
    c.email
FROM patients p
    INNER JOIN contacts c ON c.patient_id = p.id
WHERE
    p.deleted_ts IS NULL
ORDER BY p.last_name
LIMIT $1
OFFSET
    $2;