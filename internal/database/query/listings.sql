-- name: GetListingByPropertyID :one
SELECT * FROM listings
WHERE property_id = $1;

-- name: UpdateAcceptedBidIdByListingId :exec
UPDATE listings 
SET 
    accepted_bid_id = $1
WHERE id = $2;
