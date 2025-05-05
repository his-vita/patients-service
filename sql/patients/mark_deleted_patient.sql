UPDATE patients
SET
    deleted_ts = CURRENT_TIMESTAMP,
    deleted_by = $2,
    version = version + 1
WHERE
    id = $1;