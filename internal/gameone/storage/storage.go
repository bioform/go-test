package storage

import (
	"sync"

	"github.com/bioform/go-test/internal/gameone/storage/file"
)

type userStorage interface {
	WriteLastPosition(string) error
	ReadLastPosition() (*string, error)
}

var instance userStorage
var once sync.Once

// GetInstance returnt a singletone instance
// See: https://medium.com/german-gorelkin/go-singleton-f408a6c11a55
func GetInstance() userStorage {
	once.Do(func() {
		instance = new(file.FileStorage)
	})

	return instance
}
