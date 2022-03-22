package config

import (
	"database/sql"
	"sync"

	_ "github.com/jackc/pgx/v4/stdlib" // postgres driver
)

var (
	underlyingConnectedDatabase *sql.DB
)

// Postgres is the configuration object for Postgres database.
type Postgres struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
	TraceOn      bool

	lock sync.Mutex
}

// Connect connect to the database as go standard database interface.
func (m *Postgres) Connect() (db *sql.DB, err error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	// Already connected before.
	if underlyingConnectedDatabase != nil {
		return underlyingConnectedDatabase, nil
	}

	db, err = sql.Open("pgx", m.DSN)
	if err != nil {
		return db, err
	}
	db.SetMaxOpenConns(m.MaxOpenConns)
	db.SetMaxIdleConns(m.MaxIdleConns)

	if pingError := db.Ping(); pingError != nil {
		return nil, pingError
	}

	// Successfully connected.
	underlyingConnectedDatabase = db
	return db, nil
}
