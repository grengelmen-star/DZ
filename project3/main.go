package main

import (
	"project3/api"
	bin "project3/bins"
	"project3/file"
	"project3/storage"
)

func main() {
	bin.NewBinList()
	api.TestFunc()
	file.TestFileFunc()
	storage.TestStorageFunc()
}
