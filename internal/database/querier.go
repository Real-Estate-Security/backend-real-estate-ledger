// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUserByEmail(ctx context.Context, email string) (Users, error)
	GetUserByID(ctx context.Context, id int64) (Users, error)
	GetUserByUsername(ctx context.Context, username string) (Users, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]Users, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error)
}

var _ Querier = (*Queries)(nil)
