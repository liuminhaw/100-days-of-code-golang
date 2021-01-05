package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	lives := 6
	endOfGame := false
	guessResult := []string{}
	guessHistory := map[string]int8{}

	// Random choose word from wordlist
	rand.Seed(time.Now().UnixNano())
	chosenWord := wordList[rand.Intn(len(wordList))]
	wordLength := len(chosenWord)

	for i := 0; i < wordLength; i++ {
		guessResult = append(guessResult, "_")
	}

	fmt.Println(logo)
	// For debugging
	// fmt.Printf("\nChosen word: %s\n", chosenWord)
	fmt.Printf("\nGuessing result: %s\n", strings.Join(guessResult, " "))

	for endOfGame != true {
		var letterGuessed string

		fmt.Print("\nGuess a letter: ")
		fmt.Scanln(&letterGuessed)

		// Check duplicate guessing
		_, ok := guessHistory[letterGuessed]
		if ok {
			fmt.Printf("You have already guessed %s before, please make another guess.\n", letterGuessed)
			continue
		} else {
			guessHistory[letterGuessed] = 1
		}

		positions, exist := FindAllInString(chosenWord, letterGuessed)
		if exist {
			for _, pos := range positions {
				guessResult[pos] = letterGuessed
			}
			fmt.Println("\nGood guess!!!")
		} else {
			fmt.Println("\nPoor guess...")
			lives--
			fmt.Printf("\n%s\n", stages[lives])
		}
		fmt.Printf("Guessing result: %s\n", strings.Join(guessResult, " "))

		exist = ContainInSlice(guessResult, "_")
		if exist != true {
			fmt.Println("\nYOU WIN!!! :)")
			endOfGame = true
		}

		if lives == 0 {
			fmt.Println("Oops, YOU LOSE :(")
			fmt.Printf("\nThe answer is %s\n", chosenWord)
			endOfGame = true
		}
	}

	// for _, v := range stages {
	// 	fmt.Println(v)
	// }
}

func FindAllInString(s string, val string) ([]int, bool) {
	positions := []int{}

	for i := 0; i < len(s); i++ {
		if s[i:i+1] == val {
			positions = append(positions, i)
		}
	}

	if len(positions) == 0 {
		return positions, false
	}
	return positions, true
}

func ContainInSlice(slice []string, val string) bool {
	for _, v := range slice {
		if val == v {
			return true
		}
	}
	return false
}
