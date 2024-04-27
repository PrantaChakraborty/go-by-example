package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	demoMap := map[string]string{"a": "apple", "b": "ball"}
	for k, v := range demoMap {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// print only keys in map using range
	for k := range demoMap {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
		fmt.Printf("%x", c)
	}

}
