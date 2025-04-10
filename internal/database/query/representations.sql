-- name: CreateRepresentation :one
INSERT INTO representations(
    user_id,
    agent_id,
    start_date,
    end_date
) VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: AcceptRepresentation :one
UPDATE representations
SET
    status = 'accepted',
    is_active = TRUE,
    signed_at = $1
WHERE id = $2
RETURNING *;

-- name: RejectRepresentation :one
UPDATE representations
SET
    status = 'rejected',
    is_active = FALSE
WHERE id = $1
RETURNING *;

-- name: ListRepresentationsByUserID :many
SELECT 
    r.id,
    r.user_id AS client_id,
    u.first_name AS client_first_name,
    u.last_name AS client_last_name,
    u.username AS client_username,
    r.agent_id,
    a.first_name AS agent_first_name,
    a.last_name AS agent_last_name,
    a.username AS agent_username,
    r.start_date,
    r.end_date,
    r.status,
    r.requested_at,
    r.signed_at,
    r.is_active
FROM representations r
JOIN users u ON r.user_id = u.id
JOIN users a ON r.agent_id = a.id
WHERE r.user_id = $1
ORDER BY r.id
LIMIT $2
OFFSET $3;

-- name: ListRepresentationsByAgentID :many
SELECT 
    r.id,
    r.user_id AS client_id,
    u.first_name AS client_first_name,
    u.last_name AS client_last_name,
    u.username AS client_username,
    r.agent_id,
    a.first_name AS agent_first_name,
    a.last_name AS agent_last_name,
    a.username AS agent_username,
    r.start_date,
    r.end_date,
    r.status,
    r.requested_at,
    r.signed_at,
    r.is_active
FROM representations r
JOIN users u ON r.user_id = u.id
JOIN users a ON r.agent_id = a.id
WHERE r.agent_id = $1
ORDER BY r.id
LIMIT $2
OFFSET $3;

-- name: GetRepresentationByID :one
SELECT 
    r.id,
    r.user_id AS client_id,
    u.first_name AS client_first_name,
    u.last_name AS client_last_name,
    u.username AS client_username,
    r.agent_id,
    a.first_name AS agent_first_name,
    a.last_name AS agent_last_name,
    a.username AS agent_username,
    r.start_date,
    r.end_date,
    r.status,
    r.requested_at,
    r.signed_at,
    r.is_active
FROM representations r
JOIN users u ON r.user_id = u.id
JOIN users a ON r.agent_id = a.id
WHERE r.id = $1;

-- name: UpdateRepresentation :one
UPDATE representations
SET
    user_id = $1,
    agent_id = $2,
    start_date = $3,
    end_date = $4,
    status = $5,
    requested_at = $6,
    signed_at = $7,
    is_active = $8
WHERE id = $9
RETURNING *;

-- name: DeleteRepresentation :exec
DELETE FROM representations
WHERE id = $1;

