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
