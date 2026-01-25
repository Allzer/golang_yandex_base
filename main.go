package main

import "fmt"

func main() {

	a := 2
	b := 3

	var res int
	
	quatro(&a, &b, &res)

	fmt.Println(res)

}

func quatro(num_a, num_b, res *int) {
	*res = *num_a + *num_b
}