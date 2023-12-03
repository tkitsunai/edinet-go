package gateway

import (
	"bytes"
	"encoding/gob"
)

func decode[T any](data []byte) (T, error) {
	var emp T
	var result T
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&result)
	if err != nil {
		return emp, err
	}
	return result, nil
}
