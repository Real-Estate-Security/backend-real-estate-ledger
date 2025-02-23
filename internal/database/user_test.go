package database

import (
	"backend_real_estate/util"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)


func createRandomUser(t *testing.T) Users {
	
	fmt.Println("Creating a random user...")
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
	require.WithinDuration(t, arg.Dob, user.Dob, 23 * time.Hour)
	require.Equal(t, arg.Role, user.Role)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
