package database

import (
	"github.com/hashicorp/go-memdb"
	"log"
)

// ConnectToDB is a func which create the table in database and connect to it
func ConnectToDB() *memdb.MemDB {

	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"todo": &memdb.TableSchema{
				Name: "todo",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Hash"},
					},
					"name": &memdb.IndexSchema{
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"description": &memdb.IndexSchema{
						Name:         "description",
						Unique:       false,
						Indexer:      &memdb.StringFieldIndex{Field: "Description"},
						AllowMissing: true,
					},
					"status": &memdb.IndexSchema{
						Name:    "status",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Status"},
					},
					"type": &memdb.IndexSchema{
						Name:    "type",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "ItemType"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		log.Fatal(err)
	}
	return db

}