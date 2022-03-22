package models

import (
	"regexp"
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
	// RegexLogin validates User.username (borrowed from GitHub).
	// Username may only contain alphanumeric characters or single hyphens,
	// and cannot begin or end with a hyphen.
	// https://stackoverflow.com/a/58726961/1592264
	RegexLogin = regexp.MustCompile(`^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$`)
)

// User represents a person who is a pebble collector.
type User struct {
	Id        int64     `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Login     string    `db:"login" json:"login"`
	GithubId  int64     `db:"github_id" json:"github_id"`
	State     UserState `db:"state" json:"state"`
	Display   string    `db:"display" json:"display"`
	Email     string    `db:"email" json:"email,omitempty"`
	Location  string    `db:"location" json:"location"`
	Company   string    `db:"company" json:"company"`
	Avatar    string    `db:"avatar" json:"avatar"`
}

// String is a brief identifier of the user.
func (u *User) String() string {
	return "@" + u.Login
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
