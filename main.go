package main

import "fmt"

type stringMap = map[string]string

func main() {

	bookmarkMap := map[string]string{
		"bookmark#1": "Закладк №1",
		"bookmark#2": "Закладк №2",
		"bookmark#3": "Закладк №3",
	}

	var bookmarkNew string = ""

	Menu:
	for {
		param := getMenu()
		
		switch param {
		case 1:
			getAllBookmark(bookmarkMap)
		case 2:
			fmt.Println("Введите название новой закладки одним словом")
			fmt.Scan(&bookmarkNew)

			addNewBookmark(bookmarkNew, bookmarkMap)

		case 3:
			fmt.Println("Введите название закладки, которую хотите удалить")
			fmt.Scan(&bookmarkNew)

			delete(bookmarkMap, bookmarkNew)
		case 4:
			break Menu
		}
	}
}

func addNewBookmark(newName string, bookmarkMap stringMap) {
	(bookmarkMap)[fmt.Sprintf("bookmark#%d", len(bookmarkMap)+1)] = newName
	fmt.Println("Новой закладки добавлена")
}

func getAllBookmark(bookmarkMap stringMap) {
	for key, value := range(bookmarkMap){
		fmt.Println(key, value)
	}
}

func getMenu() int{
	var param int
	fmt.Println(
		`
		Выберите действие с закладками:
		- 1 Просмотр закладок
		- 2 Добавить закладку
		- 3 Удалить закладку
		- 4 Выход
		`)
	fmt.Scan(&param) 
	return param
}