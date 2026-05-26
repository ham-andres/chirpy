package auth

import (
			"time"
			"errors"

			"github.com/google/uuid" 
			"github.com/alexedwards/argon2id"
			"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (tokenTypeAccess TokenType = "chirpy-access")

func HashPassword(password string) (string, error) {
			pWHash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
			if err != nil {
				return "", err
			}
			return pWHash,nil
}

func CheckPasswordHash(password, hash string) (bool, error) {
		match, err := argon2id.ComparePasswordAndHash(password, hash)
		if err != nil {
			return false, err
		}
		return match, nil
}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
		mySigningKey := []byte(tokenSecret)
		claims := &jwt.RegisteredClaims {
				Issuer:	tokenTypeAccess,
				IssuedAt:	jwt.NewNumericDate(time.Now()),
				ExpiresAt:	jwt.NewNumericDate(time.Now().Add(expiresIn)),
				Subject:	userID.String(),
			}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)
		if err != nil {
			return "", err
		}
		return ss, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	mySigningKey := []byte(tokenSecret)
	tok, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, 
														func(token *jwt.Token)(interface{}, error)
														{return mySigningKey, nil}
													)
	if err != nil {
		return uuid.Nil, err
	}
	claims, ok := tok.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return uuid.Nil, errors.New("invalid claims")
	}
	id, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil 

}


