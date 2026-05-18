package wordgame

import (
	"fmt"
	"math/rand"
	"strings"
)

func Game(difficulty int) {
	words := []string{"Yes", "No", "Maybe", "Sharp"}
	word := strings.ToLower(words[rand.Intn(len(words))])
	shuffledWord := Shuffle(word)
	fmt.Println(shuffledWord)
	tries := 0

	for {
		answer := strings.ToLower(UserInput("Guess the Word: "))
		if answer != word {
			fmt.Println("Wrong! Try Again")
			tries++
			if tries == difficulty {
				fmt.Println("You failed!")
				break
			}
			fmt.Printf("You have %v tries left", difficulty-tries)

		} else {
			fmt.Println("Correct!")
			fmt.Println("You Won")
			break
		}
	}
}
