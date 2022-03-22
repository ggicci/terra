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
		GetUserById(int64) (*User, error)
		GetUserByLogin(string) (*User, error)
		GetUserByGithubId(int64) (*User, error)
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
		"login",
		"display",
		"email",
		"location",
		"company",
		"avatar",
	).Values(
		user.CreatedAt,
		user.UpdatedAt,
		user.GithubId,
		user.Login,
		user.Display,
		user.Email,
		user.Location,
		user.Company,
		user.Avatar,
	).Suffix("RETURNING \"id\"").RunWith(s.Master()).QueryRow().Scan(&user.Id)
}

func (s userStore) GetUserById(id int64) (user *User, err error) {
	if id <= 0 {
		return nil, nil
	}

	// if cachedUser, ok := s.cache.Users().ByGithubId(id); ok {
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

func (s userStore) GetUserByLogin(login string) (user *User, err error) {
	login = strings.TrimPrefix(login, "@")
	if login == "" {
		return nil, nil
	}

	// if cachedUser, ok := s.cache.Users().ByLogin(login); ok {
	// 	return cachedUser, nil
	// }

	user = new(User)
	s.Replica().Get(user, "SELECT * FROM users WHERE login = $1", login)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	// go s.cache.Users().CacheUser(user)
	return user, err
}

func (s userStore) GetUserByGithubId(githubId int64) (user *User, err error) {
	if githubId <= 0 {
		return nil, nil
	}

	// if cachedUser, ok := s.cache.Users().ByGithubId(githubId); ok {
	// 	return cachedUser, nil
	// }

	user = new(User)
	err = s.Replica().Get(user, "SELECT * FROM users WHERE github_id = $1", githubId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	// go s.cache.Users().CacheUser(user)
	return user, err
}
