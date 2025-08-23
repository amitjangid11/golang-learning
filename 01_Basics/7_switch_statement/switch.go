package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch
	i := 5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("other")
	}

	// Multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend")
	default:
		fmt.Println("It's Weekday")
	}

	// type switch
	whoAmI := func(k interface{}) {
		switch k.(type) { // switch j := k.(type) {
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's string")
		case bool:
			fmt.Println("It's a boolean")
		default:
			fmt.Println("Other") //fmt.Println("Other", j)
		}
	}
	whoAmI("golang")
	whoAmI(true)
	whoAmI(10.0)
}
