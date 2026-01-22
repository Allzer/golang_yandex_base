package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b float32
	fmt.Scan(&a)
	fmt.Scan(&b)

	result, error := calc(a,b)
	if error != nil{
		panic("Вы пытаетесь делить на 0")
	}
	fmt.Println("Ррезультат деления", result)
}

func calc(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("На 0 делить нельзя")
	}else{
		return a / b, nil
	}
}