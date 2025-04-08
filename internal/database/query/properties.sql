-- name: GetPropertyByID :one
SELECT * FROM properties
WHERE id = $1;

-- name: GetPropertyByAddress :one
SELECT * FROM properties
WHERE address = $1;


