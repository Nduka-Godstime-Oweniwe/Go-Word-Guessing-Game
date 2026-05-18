package wordgame

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
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
