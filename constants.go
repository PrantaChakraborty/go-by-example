package main

import (
	"fmt"
	"math"
)

const s string = "hello"

func main() {
	fmt.Println(s)
	const n = 100000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}
