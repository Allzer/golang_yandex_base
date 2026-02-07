package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/fatih/color"
)

var letterslice = []rune("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz1234567890")

type UserInfo struct {
	Login string `json:"login"`
	Psw   string `json:"psw"`
	Url   string `json:"url"`
}

func (user *UserInfo) ToByteSlice() ([]byte, error){
	file, err := json.Marshal(user)
	if err != nil{
		fmt.Println(err)
	}
	return file, nil
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

func NewUser() (*UserInfo, error) {

	loginStr := promptData("Ввведите логи")

	if loginStr == "" {
		return nil, errors.New("INCORRECT_LOGIN")
	}

	psw := promptData("Введите psw")

	urlStr := promptData("Введите url")

	user := &UserInfo{
		Login: loginStr,
		Psw:   psw,
		Url:   urlStr,
	}

	if psw == "" {
		user.genPsw(10)
	}

	return user, nil
}

func promptData(prompt string) string {
	fmt.Println(prompt)

	var res string
	fmt.Scan(&res)
	return res
}
