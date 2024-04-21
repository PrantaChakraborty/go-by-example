package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	//	using make
	s = make([]string, 7)
	fmt.Println("emp", s, "len: ", len(s), "cap: ", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "c"
	s[4] = "c"
	s[5] = "c"
	fmt.Println(s)
	s = append(s, "3")

	s = append(s, "d")
	fmt.Println(s)
	s = append(s, "e", "f")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy", c)
	fmt.Printf("type is %T", c)

	l := s[:5]
	fmt.Println("sl2", l)

	l = s[2:]
	fmt.Println("sl3", l)

	t := []string{"g", "h", "i", "j"}
	fmt.Println("dcl", t)
	fmt.Printf("type is  %T", t)
	t2 := []string{"g", "h", "i", "j"}

	if slices.Equal(t, t2) {
		fmt.Println("t = t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}

	}
	fmt.Println(twoD)
	fmt.Printf("%T", twoD)
	n := make([]string, 3)
	fmt.Println(n)
}
