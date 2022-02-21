package store

import (
	"fmt"

	"github.com/io4io/terra/store/internal"
)

type storeAdaptor func(Store) error

func (a storeAdaptor) Func(s store) error {
	return a(s)
}

// Transaction automatically commit/rollback the transaction.
func Transaction(st Store, txFunc func(Store) error) (err error) {
	return transaction(st.(store), storeAdaptor(txFunc).Func)
}

func transaction(s store, txFunc func(store) error) (err error) {
	tx, err := s.Master().(internal.SqlxTransaction).Beginx()
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			switch p := p.(type) {
			case error:
				err = p
			default:
				err = fmt.Errorf("%v", p)
			}
		}

		if err != nil {
			tx.Rollback() // NOTE: rollback error ignored
			return
		}
		err = tx.Commit()
	}()

	s.SqlStore = internal.NewSqlStoreFromTx(tx)
	return txFunc(s)
}
