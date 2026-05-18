package wordgame

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	return userInput

}

func UserOption(prompt string, limit int) int {
	for {
		str, err := strconv.Atoi(UserInput(prompt))
		if err != nil || str == 0 || str > limit {
			fmt.Println("Invalid Option")
			continue
		}

		if str <= limit {
			return str
		}
	}

}
