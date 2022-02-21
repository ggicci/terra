package models

import (
	"regexp"
	"strings"
	"time"
)

type UserState string

const (
	// UserStateNormal indicates the user is a normal user.
	UserStateNormal UserState = "normal"

	// UserStateClosed indicates the user had closed its account.
	UserStateClosed UserState = "closed"
)

var (
	// RegexUsername validates User.username (borrowed from GitHub).
	// Username may only contain alphanumeric characters or single hyphens,
	// and cannot begin or end with a hyphen.
	// https://stackoverflow.com/a/58726961/1592264
	RegexUsername = regexp.MustCompile(`^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$`)
)

// User represents a person who is a pebble collector.
type User struct {
	ID        int64     `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Username  string    `db:"username" json:"username"`
	GithubID  int64     `db:"github_id" json:"github_id"`
	State     UserState `db:"state" json:"state"`
	Display   string    `db:"display" json:"display"`
	Email     string    `db:"email" json:"email,omitempty"`
	Location  string    `db:"location" json:"location"`
	Company   string    `db:"company" json:"company"`
	Avatar    string    `db:"avatar" json:"avatar"`
}

// String returns a string representation of user id and username.
func (u *User) String() string {
	return u.Username
}

// AsPrefix returns the name in lowercase format.
func (u *User) AsPrefix() string {
	return strings.ToLower(u.Username)
}

func (u *User) PreInsert() {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
}

func (u *User) PreUpdate() {
	now := time.Now()
	u.UpdatedAt = now
}

func (u *User) Validate() error {
	return nil
}
