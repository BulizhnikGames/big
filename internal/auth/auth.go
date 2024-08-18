package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts API key from header of http request
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("incorrect authorization header format")
	}
	return vals[1], nil
}
