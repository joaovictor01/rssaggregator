package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API Key from the headers of an HTTP request
// Example:
// Authorization: ApiKey {insert api_key here}
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authentication information found")
	}

	vals := strings.Split(authHeader, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("invalid authentication header")
	}

	return vals[1], nil
}
