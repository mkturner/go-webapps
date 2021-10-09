// every program needs at least a main function
package main

import "fmt"

// func defines a function; name it main to satisfy pacakge
func main() {
	// print to the screen
	fmt.Println("Hello World")

	// declare variables
	var somethingToSay string
	var i int

	// assign a value to the variable
	somethingToSay = "Goodbye, Sayaonara, Ciao!"

	// Print the value to the screen
	fmt.Println(somethingToSay)

	i = 9
	fmt.Println("i is set to", i)

	// assign with type inference
	whatTheySaid := saySomething()
	fmt.Println("The saySomething function returned:", whatTheySaid)

	// capture multiple values returned
	firstSaid, moreSaid := sayMore()
	fmt.Println("The sayMore function returned:", firstSaid, moreSaid)
}

// Declare custom function that returns data of type string
func saySomething() string {
	return "something"
}

// Declare custom function that returns multiple values
func sayMore() (string, string) {
	return "something", "more"
}
