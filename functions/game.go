package wordgame

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Replay() int {
	fmt.Println("1. Replay")
	fmt.Println("2. Quit")
	return UserOption("Select Option: ", 2)
}

func Game(Tries int, difficulty int) {
	score := 0
	replay := 0
	for replay != 2 {
		clearScreen()
		words := []string{"Yes", "No", "Maybe", "Sharp", "Mathematics", "Learn", "Gifted", "Data", "daisy", "good"}
		word := strings.ToLower(words[rand.Intn(len(words))])
		shuffledWord := Shuffle(word)
		fmt.Println(shuffledWord)
		tries := 0

		for {
			answer := strings.ToLower(UserInput("Guess the Word: "))
			if answer != word {
				fmt.Println("Wrong! Try Again")
				time.Sleep(1 * time.Second)
				tries++
				if tries == Tries {
					fmt.Println("You failed!")
					time.Sleep(1 * time.Second)
					score = 0
					break
				}
				fmt.Printf("You have %v tries left", Tries-tries)

			} else {
				fmt.Println("Correct!")
				time.Sleep(1 * time.Second)
				fmt.Println("You Won!")
				score += difficulty
				highscore := GetHighScore()
				if score > highscore {
					NewHighScore(score)
					time.Sleep(1 * time.Second)
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
	clearScreen()
	fmt.Println("1. Easy")
	fmt.Println("2. Medium")
	fmt.Println("3. Hard")
	option := UserOption("Select Option: ", 3)
	fmt.Println("Loading....")
	time.Sleep(2 * time.Second)
	Game(7-option*2, option*2-1)
}
