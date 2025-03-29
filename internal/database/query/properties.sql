-- name: GetPropertyByID :one
SELECT * FROM properties
WHERE id = $1;
