package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

var numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

var symbols = []string{"!", "#", "$", "%", "&", "(", ")", "*", "+"}

func main() {
	rand.Seed(time.Now().Unix())

	var lettersCount, symbolsCount, numbersCount int
	var passwordSlice []string

	fmt.Println("Welcome to GoPassword Generator")
	fmt.Println("How many letters would you like in your password?")
	fmt.Scanln(&lettersCount)
	fmt.Println("How many symbols would you like?")
	fmt.Scanln(&symbolsCount)
	fmt.Println("How many numbers would you like?")
	fmt.Scanln(&numbersCount)

	for i := 0; i < lettersCount; i++ {
		passwordSlice = append(passwordSlice, letters[rand.Intn(len(letters))])
	}

	for i := 0; i < symbolsCount; i++ {
		passwordSlice = append(passwordSlice, symbols[rand.Intn(len(symbols))])
	}

	for i := 0; i < numbersCount; i++ {
		passwordSlice = append(passwordSlice, numbers[rand.Intn(len(numbers))])
	}

	fmt.Printf("Origin slice: %s\n", passwordSlice)

	rand.Shuffle(len(passwordSlice), func(i, j int) {
		passwordSlice[i], passwordSlice[j] = passwordSlice[j], passwordSlice[i]
	})
	fmt.Printf("Shuffled slice: %s\n", passwordSlice)

	password := strings.Join(passwordSlice, "")
	fmt.Printf("Your password is: %s\n", password)
}
