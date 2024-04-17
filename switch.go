// https://gobyexample.com/switch
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write ", i, "as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	t := time.Now().Weekday()
	switch t {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")

	default:
		fmt.Println("It's a weekday")
	}
	t1 := time.Now()
	switch {
	case t1.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")

		case int:
			fmt.Println("I'm an int")

		default:
			fmt.Printf("Don't know the type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")

}
