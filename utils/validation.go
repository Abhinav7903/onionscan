package utils

import (
	"regexp"
	"strings"
)

func IsOnion(identifier string) bool {
	// Updated regex to handle both 16 and 56 character onion addresses.
	matched, _ := regexp.MatchString(`(^|\.)[a-z2-7]{16,56}\.onion($|\/)`, strings.ToLower(identifier))
	return matched
}
