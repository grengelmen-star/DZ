package storage

import (
	"encoding/json"
	"fmt"
	"os"
	bin "project3/bins"
)

func SaveBinList(fileName string, bins bin.BinList) error {
	data, err := json.Marshal(bins)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return err
	}
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return err
	}
	return nil
}
func LoadBinList(fileName string) (bin.BinList, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return bin.BinList{}, nil
		}
		fmt.Println("Ошибка чтения файла:", err)
		return nil, err
	}
	var bins bin.BinList
	err = json.Unmarshal(data, &bins)
	return bins, nil
}
func TestStorageFunc() {}
