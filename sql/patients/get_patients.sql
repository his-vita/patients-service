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
    ip_oms.id,
    ip_oms.number,
    ip_oms.insurance_company_id,
    ip_dms.id,
    ip_dms.number,
    ip_dms.insurance_company_id,
    pd.id,
    pd.series,
    pd.number,
    pd.document_type_id
FROM
    patients p
    INNER JOIN contacts c ON c.patient_id = p.id
    INNER JOIN snils s ON s.patient_id = p.id
    INNER JOIN inn i ON i.patient_id = p.id
    LEFT JOIN insurance_policies ip_oms ON (
        ip_oms.patient_id = p.id
        AND ip_oms.type = 1
        AND ip_oms.main = true
    )
    LEFT JOIN insurance_policies ip_dms ON (
        ip_dms.patient_id = p.id
        AND ip_dms.type = 2
        AND ip_dms.main = true
    )
    LEFT JOIN personal_documents pd ON pd.patient_id = p.id
    AND pd.main = true
WHERE
    p.deleted_ts IS NULL
ORDER BY p.last_name
LIMIT $1
OFFSET
    $2;