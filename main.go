package main

import (
	"golang_learinig/account"
	"golang_learinig/files"
)

func main() {
	user, err := account.NewUser()
	if err != nil {
		return
	}

	userBytesSlice, err := user.ToByteSlice()

	files.WriteFile(
		files.WriteFileParams{
			Content:  userBytesSlice,
			NameFile: "userInfo.json",
		})

	// files.ReadFile(files.WriteFileParams{
	// 	NameFile: "userinfo.txt",
	// })
	user.OutputInfo()
}
