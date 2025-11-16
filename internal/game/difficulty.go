package game

import "fmt"

var highScore = map[string]int{
	"Easy":   0,
	"Medium": 0,
	"Hard":   0,
}

func SelectDifficulty() (string, int) {

	fmt.Println("\nSelect Difficulty Level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	var choice int
	fmt.Println("Enter your choice:")
	fmt.Scanln(&choice)

	switch choice {

	case 1:
		return "Easy", 10
	case 2:
		return "Medium", 5
	case 3:
		return "Hard", 3
	default :
		fmt.Println("nvalid choice! Defaulting to Medium.")
		return "Medium",  5
	}

}

func UpdatedHighScore (difficulty string , attempts int){
	if highScore[difficulty] == 0 || attempts < highScore[difficulty]{
		highScore[difficulty] = attempts
		fmt.Printf("New High Score for %s difficulty: %d attempts!\n", difficulty, attempts)
	}else{
		fmt.Printf("Your best for %s remains %d attempts.\n", difficulty, highScore[difficulty])
	}
}
