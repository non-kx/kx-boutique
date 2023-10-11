package biz

import (
	"context"
	"time"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type User struct {
	Id           string    `json:"id"`
	UserName     string    `json:"user_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserRepo interface {
	GetUserByID(context.Context, string) (*User, error)
	GetUserByName(context.Context, string) (*User, error)
	GetUserByEmail(context.Context, string) (*User, error)
	CreateUser(context.Context, *User) (*User, error)
	UpdateUser(context.Context, *User) (*User, error)
	DeleteUser(context.Context, string) error
}
