package utils

import (
	"sync"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type Blacklist struct {
	mu     sync.Mutex
	tokens map[string]struct{}
}

func NewBlacklist() *Blacklist {
	return &Blacklist{
		tokens: make(map[string]struct{}),
	}
}

// Add token to blacklist
func (b *Blacklist) Add(token string) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.tokens[token] = struct{}{}
	return nil // Assuming adding token to blacklist always succeeds
}

// Check if token is blacklisted
func (b *Blacklist) IsBlacklisted(token string) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	_, exists := b.tokens[token]
	return exists
}
