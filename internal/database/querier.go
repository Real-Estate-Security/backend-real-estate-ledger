// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"context"
)

type Querier interface {
	AcceptBid(ctx context.Context, id int64) error
	AcceptRepresentation(ctx context.Context, arg AcceptRepresentationParams) (Representations, error)
	CreateBid(ctx context.Context, arg CreateBidParams) (Bids, error)
	CreateRepresentation(ctx context.Context, arg CreateRepresentationParams) (Representations, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	DeleteRepresentation(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	GetListingByPropertyID(ctx context.Context, propertyID int64) (Listings, error)
	GetPropertyByID(ctx context.Context, id int64) (Properties, error)
	GetRepresentationByID(ctx context.Context, id int64) (GetRepresentationByIDRow, error)
	GetUserByEmail(ctx context.Context, email string) (Users, error)
	GetUserByID(ctx context.Context, id int64) (Users, error)
	GetUserByUsername(ctx context.Context, username string) (Users, error)
	ListBids(ctx context.Context, buyerID int64) ([]Bids, error)
	ListBidsOnListing(ctx context.Context, listingID int64) ([]Bids, error)
	ListRepresentationsByAgentID(ctx context.Context, arg ListRepresentationsByAgentIDParams) ([]ListRepresentationsByAgentIDRow, error)
	ListRepresentationsByUserID(ctx context.Context, arg ListRepresentationsByUserIDParams) ([]ListRepresentationsByUserIDRow, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]Users, error)
	RejectBid(ctx context.Context, id int64) error
	RejectRepresentation(ctx context.Context, id int64) (Representations, error)
	UpdateRepresentation(ctx context.Context, arg UpdateRepresentationParams) (Representations, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error)
}

var _ Querier = (*Queries)(nil)
