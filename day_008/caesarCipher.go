package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var alphabets = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

func main() {
	fmt.Println(logo)

	RecursivePlay()

	fmt.Println("Bye Bye!")
}

func RecursivePlay() {

	direction := LowercaseInput("Type 'encode' to encrypt, type 'decode' to decrypt: ")

	for direction != "encode" && direction != "decode" {
		direction = LowercaseInput("Type 'encode' to encrypt, type 'decode' to decrypt: ")
	}

	plainText := LowercaseInput("Enter your message:\n")

	var shift int
	fmt.Print("Enter shift position: ")
	fmt.Scanln(&shift)

	Caesar(direction, plainText, shift)

	again := LowercaseInput("Run caesar cipher again? (yes/no): ")
	if again == "yes" {
		RecursivePlay()
	}
}

func Caesar(method, text string, shift int) {
	sliceLen := len(alphabets)
	newString := ""

	for _, char := range text {
		pos := RunePosInSlice(alphabets, char)
		if pos != -1 {
			if method == "encode" {
				pos = mod((pos + shift), sliceLen)
			} else {
				pos = mod((pos - shift), sliceLen)
			}
			newString += string(alphabets[pos])
		} else {
			newString += string(char)
		}
	}

	fmt.Printf("The %sd text is %s\n", method, newString)
}

func LowercaseInput(message string) string {
	fmt.Print(message)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	return strings.ToLower(text)
}

func RunePosInSlice(slice []rune, val rune) int {
	for i, v := range slice {
		if val == v {
			return i
		}
	}
	return -1
}

func mod(a, b int) int {
	return (a%b + b) % b
}
