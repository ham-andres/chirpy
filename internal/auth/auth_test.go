package auth

import (
	"testing"
	"time"
	"github.com/google/uuid"
	
)

func TestHashPassword(t *testing.T) {
		password := "StancanTry@$$"
		hash, err := HashPassword(password)
		if err != nil {
			t.Fatalf("hashin failed")
		}
		result, err := CheckPasswordHash(password, hash)
		if err != nil {
			t.Fatalf("checking failed")
		}
		result2, err := CheckPasswordHash("apple", hash)
		if err != nil {
			t.Fatalf("checking failed")
		}
		if result != true {
			t.Errorf("expected true, got %v", result)
		}
		if result2 != false {
			t.Errorf("expected false, got %v", result2)
		}

		
}

func TestValidateJWT(t *testing.T) {
		userID := uuid.New()
		validToken, _ := MakeJWT(userID, "secret", time.Hour)

		tests := []struct {
			name        string
			tokenString string
			tokenSecret string
			wantUserID  uuid.UUID
			wantErr     bool
		}{
			{
				name:        "Valid token",
				tokenString: validToken,
				tokenSecret: "secret",
				wantUserID:  userID,
				wantErr:     false,
			},
			{
				name:        "Invalid token",
				tokenString: "invalid.token.string",
				tokenSecret: "secret",
				wantUserID:  uuid.Nil,
				wantErr:     true,
			},
			{
				name:        "Wrong secret",
				tokenString: validToken,
				tokenSecret: "wrong_secret",
				wantUserID:  uuid.Nil,
				wantErr:     true,
			},
		}
	
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotUserID, err := ValidateJWT(tt.tokenString, tt.tokenSecret)
				if (err != nil) != tt.wantErr {
					t.Errorf("ValidateJWT() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotUserID != tt.wantUserID {
					t.Errorf("ValidateJWT() gotUserID = %v, want %v", gotUserID, tt.wantUserID)
				}
			})
		}
	}
}
