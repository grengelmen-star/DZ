package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/url"
	"os"
	passgen "passwordManager/GeneratePassword"
	"passwordManager/accountfiles"
	"strings"
)

func main() {
	run()
}
func run() {
	accounts, err := accountfiles.LoadAccountFromFile("accounts.json")
	if err != nil {
		fmt.Println("Ошибка загрузки аккаунтов", err)
		return
	}
	for {
		choise := showMeny()
		switch choise {
		case "1":
			newAcc := createAccount()
			accounts = append(accounts, newAcc)
			err := accountfiles.SaveAccountToFile(accounts, "accounts.json")
			if err != nil {
				fmt.Println("Ошибка сохранения:", err)
			} else {
				fmt.Println("Аккаунт успешно сохранен!")
			}
		case "2":
			findAndShowAcc(accounts)
		case "3":
			accounts = deleteAccount(accounts)
			err := accountfiles.SaveAccountToFile(accounts, "accounts.json")
			if err != nil {
				fmt.Println("Ошибка сохранения: ", err)
			}
		case "4":
			fmt.Println("___Выход___")
			return
		default:
			fmt.Println("Такой функции пока не реализовано, пожалуйста выберите 1-4")
		}
	}
}
func showMeny() string {
	var userChoise string
	fmt.Println("----Добро пожаловать в менеджер аккаунтов----")
	fmt.Println("Пожалуйста, выберите интересующий вас пункт(1-4):")
	fmt.Println("1)Создать аккаунт\n2)Найти аккаунт\n3)Удалить аккаунт\n4)Выйти из программы")
	fmt.Scan(&userChoise)
	return userChoise
}
func findAndShowAcc(accounts []accountfiles.Account) {
	var login string
	fmt.Println("Введите логин для поиска:")
	fmt.Scan(&login)
	for _, acc := range accounts {
		if acc.Login == login {
			fmt.Printf("Найден: Login: %s, Password: %s, URL: %s\n", acc.Login, acc.Password, acc.URL)
			return
		}
	}
	fmt.Println("Такого аккаунта не существует ! ")

}
func deleteAccount(accounts []accountfiles.Account) []accountfiles.Account {
	var login string
	fmt.Println("Введите логин аккаунта, который хотите удалить:")
	fmt.Scan(&login)
	for i, acc := range accounts {
		if acc.Login == login {
			accounts = append(accounts[:i], accounts[i+1:]...)
			fmt.Println("Аккаунт успешно удален")
			return accounts
		}
	}
	fmt.Println("Такого аккаунта не существует")
	return accounts
}
func createAccount() accountfiles.Account {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	URL := promptData("Введите URL")
	myAccount1, err := newAccount(login, password, URL)
	if err != nil {
		fmt.Println("Ошибка создания аккаунта")
	}
	if password == "" {
		fmt.Println("Пароль был сгенерирован автоматически, поскольку вы не ввели его ")
		myAccount1.Password = passgen.GeneratePassword(12)
	}
	return *myAccount1

}

func newAccount(login, password, urlString string) (*accountfiles.Account, error) {
	if login == "" {
		return nil, errors.New("Have no login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}
	return &accountfiles.Account{
		Login:    login,
		URL:      urlString,
		Password: password,
	}, nil
}

func promptData(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	result, _ := reader.ReadString('\n')
	return strings.TrimSpace(result)
}
