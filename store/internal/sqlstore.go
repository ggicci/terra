package internal

import (
	"database/sql"
	"sync/atomic"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

// for squirrel
type (
	Execer   = sqlx.Execer
	QueryerX = sqlx.Queryer

	SqlxGetter interface {
		Get(dest interface{}, query string, args ...interface{}) error
		SqGet(dest interface{}, sqlizer sq.Sqlizer) error
	}

	SqlxSelector interface {
		Select(dest interface{}, query string, args ...interface{}) error
		SqSelect(dest interface{}, sqlizer sq.Sqlizer) error
	}

	SqlxTransaction interface {
		Beginx() (*sqlx.Tx, error)
	}
)

type SqlxRunner interface {
	Execer
	QueryerX
	SqlxGetter
	SqlxSelector
}

type SqlStore interface {
	Master() SqlxRunner
	Replica() SqlxRunner
}

type sqlStore struct {
	master   *sqlx.DB
	replicas []*sqlx.DB

	rollingCounter int64 // used by GetReplica()
}

func NewSqlStore(master *sql.DB, replicas ...*sql.DB) (SqlStore, error) {

	store := &sqlStore{
		rollingCounter: 0,
	}

	store.master = sqlx.NewDb(master, "pgx")
	store.replicas = make([]*sqlx.DB, len(replicas))
	for i, replica := range replicas {
		store.replicas[i] = sqlx.NewDb(replica, "pgx")
	}

	return store, nil
}

func (s *sqlStore) Master() SqlxRunner {
	return SquirrelSqlxDB{s.master}
}

func (s *sqlStore) Replica() SqlxRunner {
	if len(s.replicas) == 0 {
		return s.Master()
	}

	seq := atomic.AddInt64(&s.rollingCounter, 1) % int64(len(s.replicas))
	return SquirrelSqlxDB{s.replicas[seq]}
}

type sqlStoreTx struct {
	tx *sqlx.Tx
}

func (s *sqlStoreTx) Master() SqlxRunner {
	return SquirrelSqlxTX{s.tx}
}

func (s *sqlStoreTx) Replica() SqlxRunner {
	return SquirrelSqlxTX{s.tx}
}

func NewSqlStoreFromTx(tx *sqlx.Tx) SqlStore {
	return &sqlStoreTx{tx}
}

// SquirrelSqlxDB integrates squirrel (sq) with sqlx.
type SquirrelSqlxDB struct {
	*sqlx.DB
}

func (db SquirrelSqlxDB) SqGet(dest interface{}, sqlizer sq.Sqlizer) error {
	sqlstr, args, err := sqlizer.ToSql()
	if err != nil {
		return err

	}
	return db.Get(dest, sqlstr, args...)
}

func (db SquirrelSqlxDB) SqSelect(dest interface{}, sqlizer sq.Sqlizer) error {
	sqlstr, args, err := sqlizer.ToSql()
	if err != nil {
		return err

	}
	return db.Select(dest, sqlstr, args...)
}

// SquirrelSqlxTX integrates squirrel (sq) with sqlx.
type SquirrelSqlxTX struct {
	*sqlx.Tx
}

func (db SquirrelSqlxTX) SqGet(dest interface{}, sqlizer sq.Sqlizer) error {
	sqlstr, args, err := sqlizer.ToSql()
	if err != nil {
		return err

	}
	return db.Get(dest, sqlstr, args...)
}

func (db SquirrelSqlxTX) SqSelect(dest interface{}, sqlizer sq.Sqlizer) error {
	sqlstr, args, err := sqlizer.ToSql()
	if err != nil {
		return err

	}
	return db.Select(dest, sqlstr, args...)
}
