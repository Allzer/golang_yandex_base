package files

import (
	"fmt"
	"os"
)

type WriteFileParams struct {
	Content  []byte
	NameFile string
}

func WriteFile(p WriteFileParams) {
	file, err := os.Create(fmt.Sprintf("files/%s", p.NameFile))
	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write(p.Content)
	defer file.Close()
	if err != nil {
		file.Close()
		return
	}
	fmt.Println("Заптсь успешна")
	file.Close()
}

func ReadFile(p WriteFileParams) {
	file_name := fmt.Sprintf("files/%s", p.NameFile)
	data, _ := os.ReadFile(file_name)
	fmt.Println(string(data))
}