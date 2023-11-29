package datastore

type Memory struct {
	db map[string]interface{}
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) Open() error {
	m.db = make(map[string]interface{})
	return nil
}

func (m *Memory) Close() error {
	m.db = make(map[string]interface{})
	return nil
}
