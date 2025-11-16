package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bhaskar2610/number-guessing/internal/utils"
)


func PlayRound() int{
	difficulty,chances := SelectDifficulty()

	target := rand.Intn(100) + 1

	attempts := 0

	startTime := time.Now()

	for chances > 0{
		guess := utils.GetIntInput("\nEnter Your guess:")

		attempts++
		chances--

		if guess == target {
			duration := time.Since(startTime)
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts! (%.2fs)\n", attempts, duration.Seconds())
			UpdatedHighScore(difficulty, attempts)
			return target
		}else{
			if guess<target{
				fmt.Printf("Incorrect! The number is greater than %d",guess)

			}else {
				fmt.Printf("Incorrect! The number is less than %d",guess)
			}
		}

	}

	return target
}

