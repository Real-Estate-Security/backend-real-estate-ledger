// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user.sql

package database

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(
    username,
    hashed_password,
    first_name,
    last_name,
    email,
    dob,
    role
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING id, username, hashed_password, first_name, last_name, email, dob, created_at, role
`

type CreateUserParams struct {
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Dob            time.Time `json:"dob"`
	Role           UserRole  `json:"role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Dob,
		arg.Role,
	)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Dob,
		&i.CreatedAt,
		&i.Role,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getAgentByID = `-- name: GetAgentByID :one
SELECT first_name, last_name, email FROM users
WHERE UserRole = agent
`

type GetAgentByIDRow struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (q *Queries) GetAgentByID(ctx context.Context) (GetAgentByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getAgentByID)
	var i GetAgentByIDRow
	err := row.Scan(&i.FirstName, &i.LastName, &i.Email)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, hashed_password, first_name, last_name, email, dob, created_at, role FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Dob,
		&i.CreatedAt,
		&i.Role,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, hashed_password, first_name, last_name, email, dob, created_at, role FROM users
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Dob,
		&i.CreatedAt,
		&i.Role,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, hashed_password, first_name, last_name, email, dob, created_at, role FROM users
WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Dob,
		&i.CreatedAt,
		&i.Role,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, hashed_password, first_name, last_name, email, dob, created_at, role FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]Users, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Users{}
	for rows.Next() {
		var i Users
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.HashedPassword,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Dob,
			&i.CreatedAt,
			&i.Role,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
    username = $1,
    hashed_password = $2,
    first_name = $3,
    last_name = $4,
    email = $5,
    dob = $6,
    role = $7
WHERE id = $8
RETURNING id, username, hashed_password, first_name, last_name, email, dob, created_at, role
`

type UpdateUserParams struct {
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Dob            time.Time `json:"dob"`
	Role           UserRole  `json:"role"`
	ID             int64     `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Username,
		arg.HashedPassword,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Dob,
		arg.Role,
		arg.ID,
	)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Dob,
		&i.CreatedAt,
		&i.Role,
	)
	return i, err
}
