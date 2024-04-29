package main

import "fmt"

func sums(nums ...int) {
	fmt.Println("nums: ", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total is: ", total)
}

func returnString(s ...string) {
	fmt.Println(s)
}

func main() {
	sums(1, 2, 3)
	sums(1, 2, 3, 4)
	arr := []int{1, 2, 3, 4, 5}
	sums(arr...)
	s := "Hello there"
	t := "hi there"
	returnString(s, t)
}
