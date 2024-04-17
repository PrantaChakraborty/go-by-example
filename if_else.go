package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Print("7 is even\n")
	} else {
		fmt.Print("7 is odd\n")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 is event")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "num is negative")
	} else if num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
