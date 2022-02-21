package store

import (
	"database/sql"
	"strings"

	"github.com/io4io/terra/store/models"
)

const UserTable = "users"

type (
	User = models.User

	UserStore interface {
		Create(*User) error
		GetUserByID(int64) (*User, error)
		GetUserByUsername(string) (*User, error)
		GetUserByGithubID(int64) (*User, error)
	}
)

type userStore store

func (s userStore) Create(user *User) (err error) {
	user.PreInsert()

	if err := user.Validate(); err != nil {
		return err
	}
	return psql.Insert(UserTable).Columns(
		"created_at",
		"updated_at",
		"github_id",
		"username",
		"display",
		"email",
		"location",
		"company",
		"avatar",
	).Values(
		user.CreatedAt,
		user.UpdatedAt,
		user.GithubID,
		user.Username,
		user.Display,
		user.Email,
		user.Location,
		user.Company,
		user.Avatar,
	).Suffix("RETURNING \"id\"").RunWith(s.Master()).QueryRow().Scan(&user.ID)
}

func (s userStore) GetUserByID(id int64) (user *User, err error) {
	if id <= 0 {
		return nil, nil
	}

	// if cachedUser, ok := s.cache.Users().ByGithubID(id); ok {
	// 	return cachedUser, nil
	// }

	user = new(User)
	err = s.Replica().Get(user, "SELECT * FROM users WHERE id = $1", id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	// go s.cache.Users().CacheUser(user)
	return user, err
}

func (s userStore) GetUserByUsername(username string) (user *User, err error) {
	username = strings.TrimPrefix(username, "@")
	if username == "" {
		return nil, nil
	}

	// if cachedUser, ok := s.cache.Users().ByUsername(username); ok {
	// 	return cachedUser, nil
	// }

	user = new(User)
	s.Replica().Get(user, "SELECT * FROM users WHERE username = $1", username)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	// go s.cache.Users().CacheUser(user)
	return user, err
}

func (s userStore) GetUserByGithubID(githubID int64) (user *User, err error) {
	if githubID <= 0 {
		return nil, nil
	}

	// if cachedUser, ok := s.cache.Users().ByGithubID(githubID); ok {
	// 	return cachedUser, nil
	// }

	user = new(User)
	err = s.Replica().Get(user, "SELECT * FROM users WHERE github_id = $1", githubID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	// go s.cache.Users().CacheUser(user)
	return user, err
}
