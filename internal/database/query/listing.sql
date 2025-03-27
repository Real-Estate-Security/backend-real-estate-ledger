-- name: CreateListing :one
INSERT INTO listings(
    property_id,
    agent_id,
    price,
    description
) VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetListingByID :one
SELECT * FROM listings
WHERE id = $1;

-- name: ListListings :many
SELECT * FROM listings
ORDER BY id;

-- name: UpdateListingStatus :one
UPDATE listings 
SET 
    listing_status = $1
WHERE id = $2
RETURNING *;

-- name: UpdateListingPrice :one
UPDATE listings 
SET 
    price = $1
WHERE id = $2
RETURNING *;

-- name: UpdateListingAcceptedBidID :one
UPDATE listings 
SET 
    accepted_bid_id = $1
WHERE id = $2
RETURNING *;