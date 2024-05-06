package main

import "fmt"

func zeroval(ival int) {
	fmt.Println("pointer address of ival is before", &ival)
	ival = 0
	fmt.Println("pointer address of ival is after", &ival)
}

func zeroptr(iptr *int) {
	fmt.Println("pointer address of iptr is before", &iptr)
	*iptr = 5
	fmt.Println("pointer address of iptr is after", &iptr)
}

func main() {
	i := 1
	fmt.Println("pointer address of i is ", &i)
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("pointer address of i is zeroval ", &i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("pointer address of i is ", &i)
	fmt.Println("zeroptr: ", i)

}
