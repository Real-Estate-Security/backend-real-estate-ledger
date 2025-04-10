-- name: CreateBid :one
INSERT INTO bids(
    id, 
    listing_id, 
    buyer_id, 
    agent_id, 
    amount, 
    status, 
    created_at, 
    previous_bid_id
) VALUES (
    nextval('bids_id_seq'::regclass),
    $1,
    $2,
    $3,
    $4,
    'pending'::"BidStatus",
    now(),
    $5
)
RETURNING *;

-- name: ListBids :many
SELECT * FROM bids
WHERE buyer_id=$1;

-- name: ListBidsOnListing :many
SELECT * FROM bids
WHERE listing_id=$1;

-- name: RejectBid :exec
UPDATE bids
SET status = 'rejected'
WHERE id=$1;

-- name: AcceptBid :exec
UPDATE bids
SET status = 'accepted'
WHERE id = $1;