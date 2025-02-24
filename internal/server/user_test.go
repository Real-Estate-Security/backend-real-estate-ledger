package server

import (
	"backend_real_estate/internal/database"
	"backend_real_estate/util"
	"time"
)

func randomUser() database.Users {

	password := util.RandomString(6)
	hashedPassword, _ := util.HashPassword(password)

	return database.Users{
		ID:             int64(util.RandomInt(1, 1000)),
		Username:       util.RandomUsername(),
		Email:          util.RandomEmail(),
		Dob:            util.RandomDOB(),
		Role:           database.UserRole(util.RandomRole()),
		HashedPassword: hashedPassword,
		FirstName:      util.RandomString(6),
		LastName:       util.RandomString(6),
		CreatedAt:      time.Now(),
	}
}
