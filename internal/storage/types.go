package storage

import "fmt"

type Storage interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type NotFoundError struct {}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Key not found")
}
