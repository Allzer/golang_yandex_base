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
		variant := getMenu()
		switch variant {
		case 1:
			account.NewUser()
		case 2:
			url := account.PromptData("Введите url")
			users := account.NewVault().FindUserByUrl(url)
			fmt.Println(users)
		case 3:
			delUser()
		case 4:
			break Menu
		}
		
	}

}

func getMenu() int{
	var variant int
	fmt.Println("Выберите вариант")
	fmt.Println("1: Создать аккаунт")
	fmt.Println("2: Найти аккаунт")
	fmt.Println("3: Удалить аккаунт")
	fmt.Println("4: Выход")

	fmt.Scan(&variant)
	return variant
}

func delUser() {

}