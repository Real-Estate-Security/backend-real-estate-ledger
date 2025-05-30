-- name: CreateProperty :one
INSERT INTO properties(
    owner,
    address,
    city,
    state,
    zipcode,
    bedrooms,
    bathrooms
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;

-- name: ListProperties :many
SELECT * FROM properties
ORDER BY id;

-- name: UpdatePropertyOwner :one
UPDATE properties 
SET 
    owner = $1
WHERE id = $2
RETURNING *;