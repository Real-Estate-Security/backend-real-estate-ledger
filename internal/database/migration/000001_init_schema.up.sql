CREATE TYPE "UserRole" AS ENUM (
  'user',
  'agent',
  'admin'
);

CREATE TYPE "BidStatus" AS ENUM (
  'pending',
  'accepted',
  'rejected',
  'countered'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "dob" timestamp NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "role" "UserRole" NOT NULL DEFAULT 'user'
);

CREATE TABLE "properties" (
  "id" bigserial PRIMARY KEY,
  "owner" bigserial NOT NULL,
  "address" varchar NOT NULL,
  "city" varchar NOT NULL,
  "state" varchar NOT NULL,
  "zipcode" int NOT NULL,
  "bedrooms" int NOT NULL,
  "bathrooms" int NOT NULL
);

CREATE TABLE "listings" (
  "id" bigserial PRIMARY KEY,
  "property_id" bigserial NOT NULL,
  "agent_id" bigserial NOT NULL,
  "price" decimal(12,2) NOT NULL,
  "listing_status" varchar NOT NULL DEFAULT 'active',
  "listing_date" timestamptz NOT NULL DEFAULT (now()),
  "description" text,
  "accepted_bid_id" bigint DEFAULT NULL
);

CREATE TABLE "bids" (
  "id" bigserial PRIMARY KEY,
  "listing_id" bigserial NOT NULL,
  "buyer_id" bigserial NOT NULL,
  "agent_id" bigserial NOT NULL,
  "amount" decimal(12,2) NOT NULL,
  "status" "BidStatus" NOT NULL DEFAULT 'pending',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "previous_bid_id" bigserial
);

CREATE TABLE "representations" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "agent_id" bigserial NOT NULL,
  "start_date" timestamptz NOT NULL DEFAULT (now()),
  "end_date" timestamptz,
  "is_active" boolean NOT NULL DEFAULT true
);

COMMENT ON COLUMN "representations"."user_id" IS 'Buyer being represented';

COMMENT ON COLUMN "representations"."agent_id" IS 'Real estate agent representing the buyer';

COMMENT ON COLUMN "representations"."start_date" IS 'Date when representation started';

COMMENT ON COLUMN "representations"."end_date" IS 'Date when representation ended, null if still active';

COMMENT ON COLUMN "representations"."is_active" IS 'Whether the representation is currently active';

ALTER TABLE "properties" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "listings" ADD FOREIGN KEY ("property_id") REFERENCES "properties" ("id");

ALTER TABLE "listings" ADD FOREIGN KEY ("agent_id") REFERENCES "users" ("id");

ALTER TABLE "listings" ADD FOREIGN KEY ("accepted_bid_id") REFERENCES "bids" ("id");

ALTER TABLE "bids" ADD FOREIGN KEY ("listing_id") REFERENCES "listings" ("id");

ALTER TABLE "bids" ADD FOREIGN KEY ("buyer_id") REFERENCES "users" ("id");

ALTER TABLE "bids" ADD FOREIGN KEY ("agent_id") REFERENCES "users" ("id");

ALTER TABLE "bids" ADD FOREIGN KEY ("previous_bid_id") REFERENCES "bids" ("id");

ALTER TABLE "representations" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "representations" ADD FOREIGN KEY ("agent_id") REFERENCES "users" ("id");
