package middleware

import (
	"strings"
	"testing"

	"github.com/feldtsen/farrago/pkg/db"
)

func TestGenerateHashPassword_WithLengthValidation(t *testing.T) {
	tests := []struct {
		password string
		wantErr  bool
	}{
		{strings.Repeat("a", 12), false},                                // Minimum valid length
		{strings.Repeat("a", 40), false},                                // Maximum valid length
		{strings.Repeat("a", 7), true},                                  // Too short
		{strings.Repeat("a", 41), true},                                 // Too long
		{"", true},                                                      // Empty
		{strings.Repeat("a", 30) + " ", true},                           // Space
		{strings.Repeat("a", 15) + " " + strings.Repeat("b", 15), true}, // Space
		{" " + strings.Repeat("a", 30), true},                           // Space
		{strings.Repeat(" ", 30), true},                                 // Space
	}

	for _, tt := range tests {
		_, err := db.GenerateHashPassword(tt.password)
		if (err != nil) != tt.wantErr {
			t.Errorf("GenerateHashPassword(%v) error = %v, wantErr %v", tt.password, err, tt.wantErr)
		}
	}
}

/*
func TestAuthenticateUser(t *testing.T) {
	// Mock GetStoredHashForUser to return a specific hash that matches "testPassword" + PEPPER
	originalGetStoredHashForUser := GetStoredHashForUser
	GetStoredHashForUser = func(username string) (string, error) {
		if username == "testUser" {
			hash, _ := GenerateHashPassword("testPassword")
			return hash, nil
		}
		return "", nil
	}
	defer func() { GetStoredHashForUser = originalGetStoredHashForUser }()

	tests := []struct {
		username string
		password string
		want     bool
	}{
		{"testUser", "testPassword", true},
		{"testUser", "wrongPassword", false},
	}

	for _, tt := range tests {
		if got := AuthenticateUser(tt.username, tt.password); got != tt.want {
			t.Errorf("AuthenticateUser() = %v, want %v", got, tt.want)
		}
	}
}
*/
