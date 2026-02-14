package account

import (
	"fmt"
	"golang_learinig/files"
	"math/rand/v2"
	"github.com/fatih/color"
)

var letterslice = []rune("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz1234567890")
var vault = NewVault()

type UserInfo struct {
	Login string `json:"login"`
	Psw   string `json:"psw"`
	Url   string `json:"url"`
}

func (user UserInfo) OutputInfo() {
	color.Cyan("Информация про пользователя")
	fmt.Println(user)
}

func (user *UserInfo) genPsw(n int) {
	psw := make([]rune, n)

	for i := range psw {
		psw[i] = letterslice[rand.IntN(len(letterslice))]
	}

	user.Psw = string(psw)
}

func NewUser() {
	
	user := CreateUserProfile()
	vault.AddUser(*user)
	fmt.Println(vault)
	data, err := vault.ToByteSlice()

	if err != nil{
		fmt.Println("Не удалось сделать запись")
	}

	files.WriteFile(
		files.WriteFileParams{
			Content: data,
			NameFile: "userInfo.json",
		})

}

func PromptData(prompt string) string {
	fmt.Println(prompt)

	var res string
	fmt.Scan(&res)
	return res
}

func CreateUserProfile() *UserInfo{
	loginStr := PromptData("Ввведите логи")

	psw := PromptData("Введите psw")

	urlStr := PromptData("Введите url")

	user := &UserInfo{
		Login: loginStr,
		Psw:   psw,
		Url:   urlStr,
	}

	if psw == "" {
		user.genPsw(10)
	}

	return user
}
