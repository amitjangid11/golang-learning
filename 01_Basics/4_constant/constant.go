package main

import "fmt"

const age = 30 // can be declare outside the function

//  height := 6 --> not allowed shorthant declaration

func main() {
	const name string = "golang"
	// name ="javascript" --> updation not allowed

	fmt.Println(name)
	fmt.Println(age)

	// Mulitple const
	const(
		port = 5000
		host = "localhost"
	)
	// port = 5500  -->> again updation is not allowed in constant 
	fmt.Println(port,host)

}
