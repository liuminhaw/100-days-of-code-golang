package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const rock = ` 
    _______    
---'   ____)   
      (_____)  
      (_____)  
      (____)   
---.__(___)    
`

const paper = `
    _______
---'   ____)____
          ______)
          _______)
         _______)
---.__________)
`

const scissors = `
    _______
---'   ____)____
          ______)
       __________)
      (____)
---.__(___)
`

func main() {
	var yourChoice int8
	options := []string{rock, paper, scissors}

	fmt.Print("What do you choose? Type 0 for Rock, 1 for Paper or 2 for Scissors: ")
	fmt.Scanln(&yourChoice)

	if yourChoice < 0 || yourChoice > 2 {
		fmt.Println("Invalid input")
		os.Exit(0)
	}

	// Random number between 0 - 2
	rand.Seed(time.Now().UnixNano())
	machineChoice := int8(rand.Intn(3))

	fmt.Println("Your choice: ")
	fmt.Println(options[yourChoice])
	fmt.Println("Machine choice: ")
	fmt.Println(options[machineChoice])

	// Result decision
	if yourChoice == machineChoice {
		fmt.Println("It's a draw")
	} else if (yourChoice == 1 && machineChoice == 0) || (yourChoice == 2 && machineChoice == 1) ||
		(yourChoice == 0 && machineChoice == 2) {
		fmt.Println("You win")
	} else {
		fmt.Println("You lose")
	}

}
