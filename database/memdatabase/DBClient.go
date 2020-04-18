package memdatabase

import (
	"github.com/Ofla/TODO/config"
	"github.com/hashicorp/go-memdb"
	"log"
)

// ConnectToDB is a func which create the table in database and connect to it
func ConnectToDB(conf *config.Database) *memdb.MemDB {
	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			conf.Tablename: {
				Name: conf.Tablename,
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Hash"},
					},
					"name": {
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"description": {
						Name:         "description",
						Unique:       false,
						Indexer:      &memdb.StringFieldIndex{Field: "Description"},
						AllowMissing: true,
					},
					"status": {
						Name:    "status",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Status"},
					},
					"type": {
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
