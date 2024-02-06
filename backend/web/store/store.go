package store

import (
	"TP-Back-Planity/web/store/inter"
	"database/sql"
)

func NewStore(db *sql.DB) *Store {
	return &Store{
		Client:       NewClientStore(db),
		Professional: NewProfessionalStore(db),
		Admin:        NewAdminStore(db),
		Request:      NewRequestStore(db),
		Service:      NewServiceStore(db),
	}
}

type Store struct {
	Client       inter.ClientStoreInterface
	Professional inter.ProfessionalStoreInterface
	Admin        inter.AdminStoreInterface
	Request      inter.RequestStoreInterface
	Service      inter.ServiceStoreInterface
}
