package main

import "fmt"

func main() {
	var totalBill float32
	var tipPercentage, splitAmount int

	fmt.Println("Welcome to tip calculator.")

	fmt.Print("What was the total bill? $")
	fmt.Scanln(&totalBill)

	fmt.Print("What percentage tip would you like to give? 10, 12, or 15? ")
	fmt.Scanln(&tipPercentage)

	fmt.Print("How many people to split the bill? ")
	fmt.Scanln(&splitAmount)

	perPayment := (totalBill * (1 + (float32(tipPercentage) / 100))) / float32(splitAmount)

	fmt.Printf("Each person should pay: $%.2f\n", perPayment)
}
