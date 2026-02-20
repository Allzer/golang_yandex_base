package main

import (
	"fmt"
	"golang_learinig/account"
)

//Создать аккаунт
//Найти аккаунт
//Удалить аккаунт я не сделал по причине "лень"
//Выход

func main() {
Menu:
	for {
		variant := account.PromptDataTaskGeneric([]string{
			"Выберите вариант",
			"1: Создать аккаунт",
			"2: Найти аккаунт",
			"3: Выход",
		})
		switch variant {
		case "1":
			account.NewUser()
		case "2":
			url := account.PromptData("Введите url")
			users := account.NewVault().FindUserByUrl(url)
			fmt.Println(users)
		case "3":
			break Menu
		}
		
	}

}

func delUser() {

}