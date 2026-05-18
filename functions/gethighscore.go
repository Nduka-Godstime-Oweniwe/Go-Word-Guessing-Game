package wordgame

import (
	"encoding/json"
	"os"
)

func GetHighScore() int {
	data, err := os.ReadFile("highscore.json")
	var highscore int
	if err != nil {
		data, _ := json.Marshal(0)
		err = os.WriteFile("highscore.json", data, 0644)
		return 0

	} else {
		err = json.Unmarshal(data, &highscore)
	}

	return highscore
}
