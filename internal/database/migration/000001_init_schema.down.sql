-- Drop tables in reverse order of creation to avoid conflicts

DROP TABLE IF EXISTS "representations";
DROP TABLE IF EXISTS "bids";
DROP TABLE IF EXISTS "listings";
DROP TABLE IF EXISTS "properties";
DROP TABLE IF EXISTS "users";

-- Drop enums last
DROP TYPE IF EXISTS "BidStatus";
DROP TYPE IF EXISTS "UserRole";
