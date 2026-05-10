package storage

import (
	"os"
	"path/filepath"
)

type StorageManager interface {
	Save(key string, data []byte) error
	Load(key string) ([]byte, error)
	Delete(key string) error
}
type fileStorage struct {
	basePath string
}

func NewFileStorage(basePath string) (*fileStorage, error) {
	err := os.MkdirAll(basePath, 0755)
	if err != nil {
		return nil, err
	}
	return &fileStorage{basePath: basePath}, nil
}
func (f *fileStorage) Save(key string, data []byte) error {
	path := filepath.Join(f.basePath, key)
	return os.WriteFile(path, data, 0644)
}
func TestStorageFunc() {}
