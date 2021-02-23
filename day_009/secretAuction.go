package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var moreBidder string
	players := make(map[string]int)
	keepPlaying := true

	fmt.Print(logo)

	for keepPlaying {
		bids(players)

		fmt.Print("Are there any bidders? Type 'yes' or 'no': ")
		fmt.Scanln(&moreBidder)

		if strings.ToLower(moreBidder) == "yes" {
			ClearScreen()
		} else {
			keepPlaying = false
		}
	}

	bidder, bid := findHighest(players)
	fmt.Printf("The winner is %s with a bid of $%d\n", bidder, bid)

}

func bids(bidders map[string]int) {
	var name string
	var bid int

	fmt.Print("What is your name?: ")
	fmt.Scanln(&name)
	fmt.Print("What is your bid?: ")
	fmt.Scanln(&bid)

	bidders[name] = bid
}

func findHighest(bidders map[string]int) (highestBidder string, highestBid int) {
	for k, v := range bidders {
		if v > highestBid {
			highestBid = v
			highestBidder = k
		}
	}

	return highestBidder, highestBid
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
