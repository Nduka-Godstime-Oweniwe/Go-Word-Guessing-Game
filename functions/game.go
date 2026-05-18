package wordgame

import (
	"fmt"
	"math/rand"
	"strings"
)

func Replay() int {
	fmt.Println("1. Replay")
	fmt.Println("2. Quit")
	return UserOption("Select Option: ", 2)
}

func Game(difficulty int) {
	score := 0
	replay := 0
	for replay != 2 {
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
					score = 0
					break
				}
				fmt.Printf("You have %v tries left", difficulty-tries)

			} else {
				fmt.Println("Correct!")
				fmt.Println("You Won")
				score += difficulty
				highscore := GetHighScore()
				if score > highscore {
					NewHighScore(score)
					fmt.Printf("Congratulations! You've beat your previous High Score: %v\n", highscore)
					fmt.Printf("New HighScore: %v\n", score)
				}
				break

			}

		}
		replay = Replay()
	}
}

func PlayGame() {
	fmt.Println("1. Easy")
	fmt.Println("2. Medium")
	fmt.Println("3. Hard")
	option := UserOption("Select Option", 3)
	Game(7 - option*2)
}
