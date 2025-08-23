package main

import "fmt"

func main() {
	age := 18

	// if-else
	/*
		if age >= 18 {
			fmt.Println(("Person is an Adult"))
		} else {
			fmt.Println("Person is Under Age")
		}
	*/

	// if-else ladder
	if age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 {
		fmt.Println("Person is teenager")
	} else {
		fmt.Println("Person is a kid")
	}

	// if-else using logical operators
	var role = "admin"
	var hasPermissions = true

	if role == "admin" || !hasPermissions{
		fmt.Println("yes")
	}

	// Declarating variable in if-else
	if price := 200;price > 150{
		fmt.Println("Costly ",price)
	}else if price <= 150{
		fmt.Println("Not costly can be purchased ",price)
	}else{
		fmt.Println("Fair price ",price)
	}
}
/*
# Note : 
1) Go does not have ternary operator
*/
