package main

import "fmt"

func main() {
	var message = "GOlang"
	const module = 25
	var remaingmodule = 25
	fmt.Printf("Learning %v\n", message)
	fmt.Printf("We have total %v module and remaining modules %v \n", module, remaingmodule)

	var user string
	var donemodule int

	user = "ezio"
	donemodule = 3

	fmt.Printf("User %v completed %v\n", user, donemodule)
}
