package storage

import "errors"

var (
	ErrNotFound = errors.New("Not found")
)

//DB is the interface to a simple key/store value
type DB interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte) error
}
