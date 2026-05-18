package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	wordgame "wordgame/functions"
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

func main() {

	for {
		clearScreen()
		fmt.Println("WORD GUESSING GAME")
		fmt.Println("1. Play Game")
		fmt.Println("2. View HighScore")
		fmt.Println("3. How To Play")
		fmt.Println("4. Exit")
		option := wordgame.UserOption("Select An Option: ", 4)
		if option == 1 {
			wordgame.PlayGame()
		} else if option == 2 {
			fmt.Printf("High Score: %v\n", wordgame.GetHighScore())
			wordgame.UserInput("Press Enter To Continue: ")

		} else if option == 3 {
			fmt.Println("How To Play")
		} else {
			break
		}
	}
}
