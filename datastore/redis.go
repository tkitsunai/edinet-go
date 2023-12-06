package datastore

import "github.com/redis/go-redis/v9"

type Redis struct {
	rdb *redis.Client
}

func NewRedis() *Redis {
	return &Redis{}
}

func (r *Redis) Open() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	r.rdb = rdb
	return nil
}

func (r *Redis) Close() error {
	//TODO implement me
	panic("implement me")
}

func (r *Redis) GetDriver() Driver {
	//TODO implement me
	panic("implement me")
}

func (r *Redis) FindByKey(table, key string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Redis) FindAll(table string) ([][]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Redis) Update(table, key string, data interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (r *Redis) Batch(table string, dataWithKey map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}
