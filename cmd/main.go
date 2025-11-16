package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bhaskar2610/number-guessing/internal/game"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	
	fmt.Println("Welcome to guess the number game.")
	fmt.Println("I am thing a number between 1 to 100 .")
	for {
		target := game.PlayRound()
		var choice string
		fmt.Print("\nDo you want to play agian? (y/n):")

		fmt.Scanln(&choice)

		if choice != "y" && choice != "Y" {
			fmt.Printf("Thank you for playing. Guess was %d . GoodBye!!", target)
			break
		}


	}

}
