package storage

import (
	"encoding/json"
	"fmt"
	"os"
	bin "project3/bins"
)

func CreateJsonBin(fileName string, b bin.Bin) error {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Ошибка создания JSON файла")
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	return encoder.Encode(b)
}
func ReadJsonBin(fileName string) (*bin.Bin, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл не существует !")
			return nil, err
		}
		fmt.Println("Ошибка чтения JSON файла:", err)
		return nil, err
	}
	var binAcc bin.Bin
	err = json.Unmarshal(file, &binAcc)
	if err != nil {
		fmt.Println(err)
	}
	return &binAcc, nil
}
func TestStorageFunc() {}
