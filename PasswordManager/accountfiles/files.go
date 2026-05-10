package accountfiles

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	URL      string `json:"url"`
}

func SaveAccountToFile(accounts []Account, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	return encoder.Encode(accounts)
}
func LoadAccountFromFile(fileName string) ([]Account, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Account{}, nil
		}
		return nil, fmt.Errorf("Ошибка чтения файла %s: %w", fileName, err)
	}
	var accounts []Account
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		fmt.Println(err)
	}
	return accounts, nil
}
