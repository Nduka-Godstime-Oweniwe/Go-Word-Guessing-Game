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
		option := wordgame.UserOption("Select An Option: ", 2)
		if option == 2 {
			break
		} else {
			wordgame.Game(2)
		}
	}
}
