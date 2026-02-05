package files

import (
	"fmt"
	"os"
)

func WriteFile() {
	data := []byte("Hello, Go!")
	os.WriteFile("files/example.txt", data, 0644)
}

func ReadFile() {
	data, _ := os.ReadFile("files/example.txt")
	fmt.Println(string(data))
}