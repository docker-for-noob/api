package formater

import "strings"

func NormalizeEmail(email string) string {
	return strings.ToLower(strings.Trim(email, " "))
}
