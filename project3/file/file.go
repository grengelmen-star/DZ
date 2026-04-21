package file

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(fileName string) []byte {
	file, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Ошибка: файл не существует")
		}
		fmt.Println("Ошибка чтения файла:", err)
	}
	return file
}
func IsJSONFile(fileName string) bool {
	return strings.HasSuffix(fileName, ".json")
}
func TestFileFunc() {}
