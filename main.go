package main

import (
	"fmt"
	wordgame "wordgame/functions"
)

func main() {

	for {
		fmt.Println("WORD GUESSING GAME")
		fmt.Println("1. Play Game")
		fmt.Println("2. View HighScore")
		fmt.Println("3. How To Play")
		fmt.Println("4. Exit")
		option := wordgame.UserOption("Select An Option: ", 4)
		if option == 1 {
			wordgame.Game(2)
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
