package main

import (
	"golang_learinig/account"
	"golang_learinig/files"
)

//TODO
//Создать пакет, который будет работать с файлами - "files"

func main() {
	user, err := account.NewUser()
	if err != nil {
		return
	}
	files.WriteFile()
	files.ReadFile()
	user.OutputInfo()
}