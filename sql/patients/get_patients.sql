SELECT
    p.id,
    p.first_name,
    p.last_name,
    p.middle_name,
    p.birth_date,
    p.gender,
    c.phone_number,
    c.work_phone_number,
    c.email,
    s.number,
    i.number,
    ip.number,
    ip.type,
    ip.insurance_company_id
FROM patients p
    INNER JOIN contacts c ON c.patient_id = p.id
    INNER JOIN snils s ON s.patient_id = p.id
    INNER JOIN inn i ON i.patient_id = p.id
    LEFT JOIN insurance_policies ip ON ip.patient_id = p.id
    AND (
        ip.expiry_date IS NULL
        OR ip.expiry_date >= CURRENT_DATE
    )
WHERE
    p.deleted_ts IS NULL
ORDER BY p.last_name
LIMIT $1
OFFSET
    $2;