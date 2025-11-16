package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func GetIntInput(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Please enter a valid number!")
			continue
		}
		return num
	}
}