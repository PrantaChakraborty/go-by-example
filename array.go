package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)
	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])

	// array with values
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	fmt.Println(len(b))
	// the compiler will count the array element
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
	two_d := [2][3]int{
		{1, 3, 4},
		{5, 5, 6},
	}
	fmt.Println(two_d)
	for k := range 4 {
		fmt.Println(k)
	}
}
