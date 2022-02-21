package store

// Store collected the storage APIs for the app.
type Store interface {
	Users() UserStore
	// add store...
}

// domain implements Store interface.
type domain struct {
	users UserStore
	// add store implementation...
}

func (dm *domain) Users() UserStore { return dm.users }
