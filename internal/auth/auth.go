package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts API Key from Headers of HTTP Request
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("invalid first part of auth header")
	}
	return vals[1], nil
}
