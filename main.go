package main

import "fmt"

func main() {
	var name string
	var age string
	var colour string

	fmt.Print("\nPlease enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Your age: ")
	fmt.Scan(&age)

	fmt.Print("Your favourite colour: ")
	fmt.Scan(&colour)

	fmt.Println("\nHello, my name is", name, "\nI'm", age, "years old and my favourite colour is", colour)
}
