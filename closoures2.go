package main

import "fmt"

func even() func() int {
	a := 0

	return func() int {
		a += 2
		return a
	}
}

func main() {
	e := even()
	for i := 0; i < 5; i++ {
		fmt.Println(e())
	}
	f := even()
	fmt.Println(f())

}
