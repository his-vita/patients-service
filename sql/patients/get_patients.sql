SELECT *
FROM patients
WHERE
    deleted_ts ISNULL
ORDER BY last_name
LIMIT $1
OFFSET
    $2;