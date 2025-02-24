package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString generates a random string of length n
func RandomString(length int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomUsername generates a random username
func RandomUsername() string {
	return RandomString(6) + strconv.Itoa(RandomInt(10, 99))
}

// RandomEmail generates a random email
func RandomEmail() string {
	return RandomString(6) + "@" + RandomString(3) + ".com"
}

// RandomRole generates a random role
func RandomRole() string {
	roles := []string{"user", "agent"}
	n := len(roles)
	return roles[rand.Intn(n)]
}

// RandomPassword generates a random password
func RandomPassword() string {
	return RandomString(10)
}

// RandomDOB generates a random date of birth
func RandomDOB() time.Time {
	min := time.Date(1950, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2003, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}



