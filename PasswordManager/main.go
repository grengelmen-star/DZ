package main

import (
	"bufio"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"os"
	"strings"
	"time"
)

type account struct {
	login    string
	password string
	URL      string
}
type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

func (acc *account) generatePassword(n int) {
	symbols := "abcdifghjklmnopqrstuvw"
	symbolsRune := []rune(symbols)
	required := make([]rune, 0, n)
	for i := 0; i < n; i++ {
		index := randomInt(len(symbolsRune))
		required = append(required, symbolsRune[index])
	}
	acc.password = string(required)
}

func (acc account) outputPassword() {
	fmt.Println(acc)
}

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	URL := promptData("Введите URL")
	myAccount1, err := newAccount(login, password, URL)
	if err != nil {
		return
	}
	if password == "" {
		fmt.Println("Пароль был сгенерирован автоматически, поскольку вы не ввели его ")
		myAccount1.generatePassword(12)
	}
	myAccount1.outputPassword()

}
func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("Have no login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}
	newAcc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			URL:      urlString,
			login:    login,
			password: password,
		},
	}
	return newAcc, nil
}
func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("Have no login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}
	return &account{
		login:    login,
		URL:      urlString,
		password: password,
	}, nil
}

func promptData(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	result, _ := reader.ReadString('\n')
	return strings.TrimSpace(result)
}

func randomInt(n int) int {
	if n <= 0 {
		return 0
	}
	bigN := big.NewInt(int64(n))
	val, err := rand.Int(rand.Reader, bigN)
	if err != nil {
		panic(err)
	}
	return int(val.Int64())
}
