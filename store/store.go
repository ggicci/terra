package store

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/io4io/terra/store/internal"
)

var (
	// globalStore Store
	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

// store implements Store interface.
type store struct {
	internal.SqlStore
	Store
}

func NewStore(sqlstore internal.SqlStore) Store {
	st := store{
		SqlStore: sqlstore,
	}
	domainProvider := &domain{
		users: userStore(st),
	}
	st.Store = domainProvider
	return st
}
