package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
	fmt.Println("end of for")
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
	for k := range 7 {
		print(k, "\n")
	}
	for {
		fmt.Println("loop")
		break
	}
	for i = 0; i < 10; i++ {
		print(i)
	}
}
