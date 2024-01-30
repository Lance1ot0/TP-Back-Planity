package store

import (
	"TP-Back-Planity/web/store/inter"
	"database/sql"
)

func NewStore(db *sql.DB) *Store {
	return &Store{
		Client: NewClientStore(db), // Replace 'db' with the actual *sql.DB instance
	}
}

type Store struct {
	Client inter.ClientStoreInterface
}
