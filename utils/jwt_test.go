package utils

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "testpassword"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if hashedPassword == "" {
		t.Fatal("Expected hashed password to not be empty")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		t.Fatalf("Expected hash and password to match, but got error: %v", err)
	}
}

func TestCheckPassHash(t *testing.T) {
	password := "testpassword"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if !CheckPassHash(password, hashedPassword) {
		t.Fatal("Expected password to match the hash, but it did not")
	}

	wrongPassword := "wrongpassword"
	if CheckPassHash(wrongPassword, hashedPassword) {
		t.Fatal("Expected wrong password to not match the hash, but it did")
	}
}
