package datastore

import (
	"fmt"
	"github.com/boltdb/bolt"
	"os/user"
	"path/filepath"
	"time"
)

func NewBoltDB() *BoltDB {
	return &BoltDB{db: nil}
}

type BoltDB struct {
	db *bolt.DB
}

func (b *BoltDB) Open() error {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user information:", err)
		return err
	}
	dbPath := filepath.Join(usr.HomeDir, ".edinet-go", "boltdb.db")
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		return err
	}
	b.db = db
	return nil
}

func (b *BoltDB) Close() error {
	return b.db.Close()
}
