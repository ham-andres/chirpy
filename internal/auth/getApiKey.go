package auth

import (
			"net/http"
			"errors"
			"strings"
)

//ErrNoAuthHeaderIncluded is global only for auth package so no need to be declared twice.

func GetAPIKey(headers	http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
			return "", ErrNoAuthHeaderIncluded
	}

	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
			return "", errors.New("wrong command, usage: ApiKey <api string>")
	}

	return splitAuth[1], nil 
}


