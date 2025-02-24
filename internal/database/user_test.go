package database

import (
	"backend_real_estate/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) Users {
	arg := CreateUserParams{
		Username:       util.RandomUsername(),
		HashedPassword: util.RandomPassword(),
		FirstName:      util.RandomString(6),
		LastName:       util.RandomString(6),
		Email:          util.RandomEmail(),
		Dob:            util.RandomDOB(),
		Role:           "user",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Email, user.Email)
	require.WithinDuration(t, arg.Dob, user.Dob, 23*time.Hour)
	require.Equal(t, arg.Role, user.Role)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserByID(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByID(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.Dob, user2.Dob, 23*time.Hour)
	require.Equal(t, user1.Role, user2.Role)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestGetUserByUsername(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByUsername(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.Dob, user2.Dob, 23*time.Hour)
	require.Equal(t, user1.Role, user2.Role)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByEmail(context.Background(), user1.Email)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.Dob, user2.Dob, 23*time.Hour)
	require.Equal(t, user1.Role, user2.Role)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID:             user1.ID,
		Username:       util.RandomUsername(),
		HashedPassword: util.RandomPassword(),
		FirstName:      util.RandomString(6),
		LastName:       util.RandomString(6),
		Email:          util.RandomEmail(),
		Dob:            user1.Dob,
		Role:           user1.Role,
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.HashedPassword, user2.HashedPassword)
	require.Equal(t, arg.FirstName, user2.FirstName)
	require.Equal(t, arg.LastName, user2.LastName)
	require.Equal(t, arg.Email, user2.Email)
	require.WithinDuration(t, user1.Dob, user2.Dob, 23*time.Hour)
	require.Equal(t, user1.Role, user2.Role)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user1.ID)

	require.NoError(t, err)

	user2, err := testQueries.GetUserByID(context.Background(), user1.ID)

	require.Error(t, err)
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
