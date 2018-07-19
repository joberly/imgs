package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// UserMutableData contains mutable user data for a user account. This
// data is updated as-is when using the Update on the UserDB interface.
type UserMutableData struct {
	Name  string
	Email string
}

// User contains all user data for a user account. This data is managed
// internally by each specific UserService.
type User struct {
	UUID            uuid.UUID
	CreatedAt       time.Time
	ModifiedAt      time.Time
	DeletedAt       time.Time
	UserMutableData `bson:",inline"`
}

// UserDB accesses users in persistent storage.
type UserDB interface {
	ByUUID(uuid uuid.UUID) (*User, error)
	ByEmail(email string) (*User, error)

	Create(user *User) error
	Update(user *User) error
	Delete(uuid uuid.UUID) error

	Close() error
}

var (
	// ErrUserNotFound indicates that the user is not in the database.
	ErrUserNotFound = errors.New("models: user not found")
)

// UserService serves users stored in persistent storage.
type UserService interface {
	UserDB
}

var _ UserService = &userService{}

type userService struct {
	UserDB
}
