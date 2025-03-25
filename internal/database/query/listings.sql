-- name: GetListingByPropertyID :one
SELECT * FROM listings
WHERE property_id = $1;
