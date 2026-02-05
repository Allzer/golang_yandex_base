package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
	"github.com/fatih/color"
)

var letterslice = []rune("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz1234567890")

type UserInfo struct {
	login, psw, url string
}

type userInfoWithTimeStamp struct {
	ceratedAt time.Time
	updatedAt time.Time
	UserInfo
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

	user.psw = string(psw)
}

func NewUser() (*UserInfo, error) {

	loginStr := promptData("Ввведите логи")

	if loginStr == "" {
		return nil, errors.New("INCORRECT_LOGIN")
	}

	psw := promptData("Введите psw")

	urlStr := promptData("Введите url")
	_, err := url.ParseRequestURI(urlStr)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("INCORRECT_URL")
	}

	user := &UserInfo{
		login: loginStr,
		psw:   psw,
		url:   urlStr,
	}

	if psw == "" {
		user.genPsw(10)
	}

	return user, nil
}

func promptData(prompt string) string{
	fmt.Println(prompt)

	var res string
	fmt.Scan(&res)
	return res
}