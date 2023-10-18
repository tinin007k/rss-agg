package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrorsAuthNotIncluded = errors.New("Auth header not included")

func GetAPIKey(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return "", ErrorsAuthNotIncluded
	}
	authString := strings.Split(authHeader, " ")
	if len(authString) < 2 || authString[0] != "ApiKey" {
		return "", errors.New("Malformed auth header")
	}
	return authString[1], nil
}
