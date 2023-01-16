package models

import (
	"database/sql"
	"fmt"

	"github.com/s33how/lenslocked/rand"
)

type Session struct {
	ID        int
	UserID    int
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

// Create will create a new session for the user provided. The session token
// will be returned as the Token field on the Session type, but only the hashed
// session token is stored in the database.
func (ss *SessionService) Create(userID int) (*Session, error) {
	token, err := rand.SessionToken()
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	// TODO: Hash the session token
	session := Session{
		UserID: userID,
		Token:  token,
		// TODO: Set the TokenHash
	}
	// TODO: Store the session in the database
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	return nil, nil
}
