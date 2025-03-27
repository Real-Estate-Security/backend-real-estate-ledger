-- name: CreateRepresentation :one
INSERT INTO representations(
    user_id,
    agent_id,
    start_date,
    end_date,
    is_active
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: AcceptRepresentation :one
UPDATE representations
SET
    status = 'accepted',
    is_active = TRUE,
    signed_date = $1
WHERE id = $2
RETURNING *;

-- name: RejectRepresentation :one
UPDATE representations
SET
    status = 'rejected',
    is_active = FALSE,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: ListRepresentationsByUserID :many
SELECT * FROM representations
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListRepresentationsByAgentID :many
SELECT * FROM representations
WHERE agent_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: GetRepresentationByID :one
SELECT * FROM representations
WHERE id = $1;

-- name: UpdateRepresentation :one
UPDATE representations
SET
    user_id = $1,
    agent_id = $2,
    start_date = $3,
    end_date = $4,
    is_active = $5
WHERE id = $6
RETURNING *;

-- name: DeleteRepresentation :exec
DELETE FROM representations
WHERE id = $1;

