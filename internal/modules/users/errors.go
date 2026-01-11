package users

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrEmailExists      = errors.New("email already exists")
	ErrInvalidUserID    = errors.New("invalid user id")
)
