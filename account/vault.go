package account

import (
	"encoding/json"
	"fmt"
	"golang_learinig/files"
	"strings"
)

type Vault struct {
	User []UserInfo `json:"accounts"`
}

type DB interface{ //интерфейс - список методов, которые должны быть в структуре
	Read() ([]byte, error)
	Write([]byte) 
}

func NewVault() *Vault {

	file, err := files.ReadFile(files.ReadFileFileParams{
		NameFile: "userInfo.json",
	})
	if err != nil {
		return &Vault{
			User: []UserInfo{},
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil{
		fmt.Println("Не удалось преобразовать в json")
	}
	
	return &vault
}

func (vault *Vault) AddUser(user UserInfo){
	vault.User = append(vault.User, user)
}

func (vault *Vault) ToByteSlice() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		fmt.Println(err)
	}
	return file, nil
}

func (vault *Vault) FindUserByUrl(url string) []UserInfo {

	users := []UserInfo{}

	for _, user := range vault.User {
		isMatched := strings.Contains(user.Url, url) //записать в Notions

		if isMatched{
			users = append(users, user)
		}
	}

	return users
}
