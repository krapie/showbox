package database

import "github.com/hashicorp/go-memdb"

var schema = &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"show": &memdb.TableSchema{
			Name: "show",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.IntFieldIndex{Field: "ID"},
				},
			},
		},
		"ticket": &memdb.TableSchema{
			Name: "ticket",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.IntFieldIndex{Field: "ID"},
				},
				"show_id": &memdb.IndexSchema{
					Name:    "show_id",
					Unique:  false,
					Indexer: &memdb.IntFieldIndex{Field: "ShowID"},
				},
				"user_id": &memdb.IndexSchema{
					Name:    "user_id",
					Unique:  false,
					Indexer: &memdb.IntFieldIndex{Field: "UserID"},
				},
			},
		},
	},
}
