package main

import "fmt"

// Arrays : Numbered sequence of specific length
func main() {
	// Note : In case of default values int -> 0,string -> "",bool -> false
	var nums [4]int

	// Adding elements to array>
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3

	// len() Function to get length of array
	fmt.Println(len(nums))

	// Printing array element one by one
	fmt.Println(nums[0])
	fmt.Println(nums[1])

	// Printing all at one time
	fmt.Println(nums)



	// New Array to show default value is 0 int his case its false
	var vals [4]bool
	vals[0] = true
	vals[2] = true
	fmt.Println(vals)


	// String array
	var name [3]string
	name[0]="goLang"
	name[1] = "GoLingo"
	fmt.Println(name)


	// Initialise and declared in single line
	number := [3]int{1,2,3}
	fmt.Println(number)


	// 2d Arrays
	numbers :=[2][2]int{{1,2},{3,4}}
	fmt.Println(numbers)
}
