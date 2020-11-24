package main

import "fmt"

func main() {
	var city, pet string

	// Greetings
	fmt.Println("Hello, this is Band Generator")

	// Ask for city
	fmt.Println("Which city do you grew up in?")
	fmt.Scanln(&city)

	// Ask for pet
	fmt.Println("What is the name of your pet?")
	fmt.Scanln(&pet)

	// Generate band name
	fmt.Printf("Band name: %s %s\n", city, pet)
}
