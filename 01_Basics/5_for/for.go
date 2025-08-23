package main

import "fmt"

//for -> only construct in go for looping

func main() {
	//while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	//infinite loop
/*	for {
		println("1")
	}
*/
	// classic for loop

	for j := 0; j < 3; j++ {
		//break
		if j == 2 {
			continue
		}

		fmt.Println(j)
	}

	// range
	for k:= range 10{
		fmt.Println(k)
	}
}
