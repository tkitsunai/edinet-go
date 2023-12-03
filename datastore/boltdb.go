package datastore

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/hashicorp/go-multierror"
	"github.com/tkitsunai/edinet-go/logger"
	"os/user"
	"path/filepath"
	"sync"
	"time"
)

const (
	MetaDataTable = "metadata"
	CompanyTable  = "company"
)

type BoltDB struct {
	db *bolt.DB
}

func NewBoltDB() *BoltDB {
	return &BoltDB{db: nil}
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

func (b *BoltDB) GetDriver() Driver {
	return b
}

func (b *BoltDB) FindByKey(bucketKey, key string) ([]byte, error) {
	if len(key) == 0 {
		return nil, nil
	}
	var result []byte
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketKey))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}

		keyData := bucket.Get([]byte(key))
		if keyData == nil {
			return fmt.Errorf(fmt.Sprintf("data not found key: %s", key))
		}
		result = keyData
		return nil
	})
	return result, err
}

func (b *BoltDB) FindAll(bucketKey string) ([][]byte, error) {
	logger.Logger.Info().Msgf("bucketKey: %s", bucketKey)
	var resultsLock sync.Mutex
	results := make([][]byte, 0)
	resultsLock.Lock()
	defer resultsLock.Unlock()

	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketKey))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}

		err := bucket.ForEach(func(k, v []byte) error {
			results = append(results, v)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (b *BoltDB) Update(bucketKey, key string, data interface{}) error {
	err := b.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketKey))
		if err != nil {
			return err
		}
		encodedData, err := encode(data)
		err = bucketPut(bucket, []byte(key), encodedData)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (b *BoltDB) Batch(bucketKey string, dataWithKey map[string]interface{}) error {
	var wg sync.WaitGroup
	var batchErr error
	for key, data := range dataWithKey {
		wg.Add(1)
		go func(key string, data interface{}) {
			defer wg.Done()
			err := b.db.Batch(func(innerTx *bolt.Tx) error {
				bucket, err := innerTx.CreateBucketIfNotExists([]byte(bucketKey))
				if err != nil {
					return err
				}
				encodedData, err := encode(data)
				if err != nil {
					return err
				}
				err = bucketPut(bucket, []byte(key), encodedData)
				return err
			})
			if err != nil {
				batchErr = multierror.Append(batchErr, fmt.Errorf("store failed, key: %s, error: %s", key, err.Error()))
			}
		}(key, data)
	}
	wg.Wait()

	return batchErr
}

func bucketPut(bucket *bolt.Bucket, key []byte, data []byte) error {
	return bucket.Put(key, data)
}

func encode(data interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
