package auth 

import (
		"errors"
		"net/http"
		"strings"
		"crypto/rand"
		"encoding/hex"
)
var ErrNoAuthHeaderIncluded = errors.New("No auth header included in request")

func GetBearerToken( headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	/* my attempt
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("wrong authorization header, usage: Bearer <token string>")
	}
	splitAuth = strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	if token == "" {
		return "", errors.New("usage: Bearer <token string>")
	}
	*/ 
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "Bearer" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil 

}

func MakeRefreshToken() string {
		key := make([]byte, 32)
		rand.Read(key) 
		refreshToken := hex.EncodeToString(key)
		return refreshToken

}
