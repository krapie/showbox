package database

import "github.com/hashicorp/go-memdb"

type DB struct {
	db *memdb.MemDB
}

func New() (*DB, error) {
	memDB, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	return &DB{
		db: memDB,
	}, nil
}
