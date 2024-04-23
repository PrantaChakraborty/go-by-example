package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println(m)

	//	get the specific key value
	v1 := m["k1"]
	fmt.Println(v1)
	//	if any key is not present in map, it will return zero
	v2 := m["k5"]
	fmt.Println(v2)

	fmt.Println(len(m))

	//	delete key, value from map
	delete(m, "k2")

	clear(m) // clear all key value

	// it returns bool if value is present or not
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	//	declare with initial value
	newMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", newMap)

	//	 map compare
	newMap2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(newMap, newMap2) {
		fmt.Println("both are equal")
	}
}
