package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// custom errors returned by VerifyToken
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// Payload contains payload data of the token
type Payload struct {
	ID uuid.UUID		`json:"id"`
	Username string		`json:"username"`
	IssuedAt time.Time	`json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a username/duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{ 
		ID: tokenID,
		Username: username,
		IssuedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Checks if token is valid
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}