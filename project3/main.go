package main

import (
	"project3/api"
	bin "project3/bins"
	"project3/file"
	"project3/storage"
	"time"
)

func main() {
	bin.NewBinList()
	api.TestFunc()
	file.TestFileFunc()
	storage.TestStorageFunc()
}

type BinManager interface {
	Create(id, name string, private bool) (*bin.Bin, error)
	GetBin(id string) (bin.Bin, error)
}
type StorageManager interface {
	Save(key string, data []byte) error
	Load(key string) ([]byte, error)
}
type binService struct {
	storage storage.Storage
}

func (s *binService) Create(id, name string, private bool) (*bin.Bin, error) {
	newBin := &bin.Bin{
		ID:        id,
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
	}
}
