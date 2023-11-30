package datastore

type Memory struct {
	db map[string]interface{}
}

func (m *Memory) Batch(table string, dataWithKey map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m *Memory) View(table, key string) ([][]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Memory) Update(table, key string, data interface{}) error {
	//TODO implement me
	panic("implement me")
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

func (m *Memory) GetDriver() Driver {
	return m
}
