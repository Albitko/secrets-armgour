package entity

import (
	"errors"
)

var (
	// ErrLoginAlreadyInUse - error for database when user already registered
	ErrLoginAlreadyInUse = errors.New("login already exists")
	// ErrInvalidCredentials - error for database when user pass invalid credentials
	ErrInvalidCredentials = errors.New("invalid credentials")
)
