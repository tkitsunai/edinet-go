package datastore

type Engine interface {
	Open() error
	Close() error
}

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
