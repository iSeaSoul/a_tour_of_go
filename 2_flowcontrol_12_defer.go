package main

import "fmt"

func main() {
	defer fmt.Println("www")
	defer fmt.Println("world")

	fmt.Println("hello")
}

// The deferred call's arguments are evaluated immediately, but the function
// call is not executed until the surrounding function returns.
// its deferred calls are executed in last-in-first-out order.