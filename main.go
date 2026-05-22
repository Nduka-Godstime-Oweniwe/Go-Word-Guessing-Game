package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
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

func Contains(words []string, word string) bool {
	for i := 0; i < len(words); i++ {
		if strings.ToUpper(word) == strings.ToUpper(words[i]) {
			return true
		}
	}
	return false
}
func main() {

	for {
		clearScreen()
		fmt.Println("WORD GUESSING GAME")
		fmt.Println("1. Play Game")
		fmt.Println("2. Upload More Words")
		fmt.Println("3. View HighScore")
		fmt.Println("4. How To Play")
		fmt.Println("5. Exit")
		option := wordgame.UserOption("Select An Option: ", 5)
		if option == 1 {
			fmt.Println("Loading Game....")
			time.Sleep(1 * time.Second)
			fmt.Println("Pls wait a little While...")
			time.Sleep(1 * time.Second)
			words := wordgame.GetWords()
			wordgame.PlayGame(words)
		} else if option == 2 {
			user := ""
			words := wordgame.GetWords()
			for {
				user = wordgame.UserInput("Enter Word: ")
				if user == "" {
					wordgame.Dumpword(words)
					break
				} else if Contains(words, user) {
					fmt.Printf("The Word \"%v\" already exist in our database! Pls type in another word\n", strings.ToLower(user))

				} else {
					words = append(words, user)
				}

			}

		} else if option == 3 {
			fmt.Printf("High Score: %v\n", wordgame.GetHighScore())
			wordgame.UserInput("Press Enter To Continue: ")

		} else if option == 4 {
			fmt.Println("How To Play")
			fmt.Println("1. A word would be shown on the screen")
			fmt.Println("2. The letters of the word would be rearranged")
			fmt.Println("3. You are to guess the word")
			fmt.Println("4. The number of tries you have depends on the difficulty level")
			fmt.Println("\tA. Easy >> 5 tries\n\tB. Medium >> 3 tries\n\tC. Hard >> 1 try only")
			fmt.Println("5. The points you get if you guess correctly depends on the difficulty level")
			fmt.Println("\tA. Easy >> 1 point\n\tB. Medium >> 3 points\n\tC. Hard >> 5 points")
			wordgame.UserInput("Press Enter To Continue: ")
		} else {
			break
		}
	}
}
