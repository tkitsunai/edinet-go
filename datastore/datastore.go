package datastore

import (
	"errors"
)

type Engine interface {
	Open() error
	Close() error
	GetDriver() Driver
}

// Driver interfaces for persistent storage
type Driver interface {
	FindByKey(table, key string) ([]byte, error)
	FindAll(table string) ([][]byte, error)
	Update(table, key string, data interface{}) error
	Batch(table string, dataWithKey map[string]interface{}) error
}

var (
	RecordNotFound = errors.New("record not found")
	BucketNotFound = errors.New("bucket/table not found")
	RequiredKey    = errors.New("must be require key")
)

var DefaultEngine = NewMemory()

func GetEngineByName(name string) Engine {
	switch name {
	case "boltdb":
		return NewBoltDB()
	case "memory":
		return DefaultEngine
	}
	return DefaultEngine
}
