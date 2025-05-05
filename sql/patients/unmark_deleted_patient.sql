UPDATE patients
SET
    deleted_ts = NULL,
    deleted_by = NULL,
    updated_ts = CURRENT_TIMESTAMP,
    updated_by = $2,
    version = version + 1
WHERE
    id = $1;