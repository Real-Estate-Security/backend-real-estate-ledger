-- name: CreateUser :one
INSERT INTO users(
    username,
    hashed_password,
    first_name,
    last_name,
    email,
    dob,
    role
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

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetAgentByID :one
SELECT first_name, last_name, email FROM users
WHERE UserRole = agent;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET
    username = $1,
    hashed_password = $2,
    first_name = $3,
    last_name = $4,
    email = $5,
    dob = $6,
    role = $7
WHERE id = $8
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

