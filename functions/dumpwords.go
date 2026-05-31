package wordgame

import (
	"encoding/json"
	"os"
)

func Dumpword(words []string) {
	data, _ := json.Marshal(words)
	os.WriteFile("words.json", data, 0644)
}

func GetWords() []string {
	data, err := os.ReadFile("words.json")
	var words []string
	if err != nil {
		data, _ := json.Marshal([]string{})
		os.WriteFile("words.json", data, 0644)
		return words
	} else {
		err = json.Unmarshal(data, &words)
	}
	return words

}
