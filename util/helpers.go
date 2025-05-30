package util

import (
	"regexp"
)

// IsEmail checks if a given string is a valid email address.
func IsEmail(input string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(input)
}
