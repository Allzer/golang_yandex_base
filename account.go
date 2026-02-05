package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterslice = []rune("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz1234567890")

type userInfo struct {
	login, psw, url string
}

type userInfoWithTimeStamp struct {
	ceratedAt time.Time
	updatedAt time.Time
	userInfo
}

func (user userInfo) outputInfo() {
	fmt.Println(user)
}

func (user *userInfo) genPsw(n int) {
	psw := make([]rune, n)

	for i := range psw {
		psw[i] = letterslice[rand.IntN(len(letterslice))]
	}

	user.psw = string(psw)
}

func newUser() (*userInfo, error) {

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

	user := &userInfo{
		login: loginStr,
		psw:   psw,
		url:   urlStr,
	}

	if psw == "" {
		user.genPsw(10)
	}

	return user, nil
}