package auth

import (
	"testing"
	
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
