package main

import (
	"fmt"
)

func main() {
	array := []int{}

	enterTransactions(&array)
	fmt.Print(array)
}

func enterTransactions(array *[]int) {
	transaction := 0
	for {
		fmt.Scan(&transaction)
		if transaction == -10 || transaction == 10 || transaction == 40 {
			*array = append(*array, transaction)
		}
		if len(*array) == 3{
			break
		}
	} 
}