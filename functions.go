package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := sum(4, 5)
	fmt.Println(res)

	res = plusPlus(3, 4, 6)
	fmt.Println(res)
}
