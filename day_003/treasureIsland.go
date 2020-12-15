package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var choice1, choice2, choice3 string

	fmt.Println("Welcome to treasure island, your mission is to stay alive.")

	fmt.Print("Left or right: ")
	fmt.Scanln(&choice1)
	if strings.ToLower(choice1) == "right" {
		fmt.Println("Ah ha! Game Over!")
		os.Exit(0)
	}
	if strings.ToLower(choice1) != "left" {
		fmt.Println("Invalid input! Game Over!")
		os.Exit(0)
	}

	fmt.Print("Wait or swim: ")
	fmt.Scanln(&choice2)
	if strings.ToLower(choice2) == "swim" {
		fmt.Println("Ah ha! Game Over!")
		os.Exit(0)
	}
	if strings.ToLower(choice2) != "wait" {
		fmt.Println("Invalid input! Game Over!")
		os.Exit(0)
	}

	fmt.Print("Pick a color - red, blue or yellow: ")
	fmt.Scanln(&choice3)
	if strings.ToLower(choice3) == "blue" {
		fmt.Println("You did it!!!")
		os.Exit(0)
	} else if strings.ToLower(choice3) == "yellow" || strings.ToLower(choice3) == "red" {
		fmt.Println("Ah ha! Game Over!")
		os.Exit(0)
	} else {
		fmt.Println("Invalid input! Game Over!")
	}
}
