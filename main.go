package main

import (
	"fmt"
)

func main() {
	user, err := newUser()
	if err != nil {
		return
	}
	user.outputInfo()
}

func promptData(prompt string) string{
	fmt.Println(prompt)

	var res string
	fmt.Scan(&res)
	return res
}