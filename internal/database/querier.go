// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"context"
)

type Querier interface {
	AcceptRepresentation(ctx context.Context, arg AcceptRepresentationParams) (Representations, error)
	CreateListing(ctx context.Context, arg CreateListingParams) (Listings, error)
	CreateProperty(ctx context.Context, arg CreatePropertyParams) (Properties, error)
	CreateRepresentation(ctx context.Context, arg CreateRepresentationParams) (Representations, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	DeleteRepresentation(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	GetAgentByID(ctx context.Context) (GetAgentByIDRow, error)
	GetListingByID(ctx context.Context, id int64) (Listings, error)
	GetListingByPropertyID(ctx context.Context, propertyID int64) (Listings, error)
	GetListings(ctx context.Context) ([]GetListingsRow, error)
	GetPropertyByID(ctx context.Context, id int64) (Properties, error)
	GetRepresentationByID(ctx context.Context, id int64) (GetRepresentationByIDRow, error)
	GetUserByEmail(ctx context.Context, email string) (Users, error)
	GetUserByID(ctx context.Context, id int64) (Users, error)
	GetUserByUsername(ctx context.Context, username string) (Users, error)
	ListListings(ctx context.Context) ([]Listings, error)
	ListProperties(ctx context.Context) ([]Properties, error)
	ListRepresentationsByAgentID(ctx context.Context, arg ListRepresentationsByAgentIDParams) ([]ListRepresentationsByAgentIDRow, error)
	ListRepresentationsByUserID(ctx context.Context, arg ListRepresentationsByUserIDParams) ([]ListRepresentationsByUserIDRow, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]Users, error)
	RejectRepresentation(ctx context.Context, id int64) (Representations, error)
	UpdateListingAcceptedBidID(ctx context.Context, arg UpdateListingAcceptedBidIDParams) (Listings, error)
	UpdateListingPrice(ctx context.Context, arg UpdateListingPriceParams) (Listings, error)
	UpdateListingStatus(ctx context.Context, arg UpdateListingStatusParams) (Listings, error)
	UpdatePropertyOwner(ctx context.Context, arg UpdatePropertyOwnerParams) (Properties, error)
	UpdateRepresentation(ctx context.Context, arg UpdateRepresentationParams) (Representations, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error)
}

var _ Querier = (*Queries)(nil)
