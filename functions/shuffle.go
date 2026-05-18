package wordgame

import (
	"math/rand"
)

func Contains(slice []int, num int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == num {
			return true
		}
	}
	return false
}

func ShuffleIndex(slice []int) []int {
	result := []int{}
	num := 0
	for len(result) != len(slice) {
		num = rand.Intn(len(slice))
		if !Contains(result, num) {
			result = append(result, num)

		}
	}
	return result
}

func SliceOfIndex(str string) []int {
	result := []int{}
	for i, _ := range str {
		result = append(result, i)

	}
	return result
}

func StrToSlice(str string) []string {
	result := []string{}
	for _, v := range str {
		result = append(result, string(v))
	}
	return result
}
func Shuffle(str string) string {
	result := ""
	for {
		slice := StrToSlice(str)
		sliceOfIndex := SliceOfIndex(str)
		shuffledSlice := ShuffleIndex(sliceOfIndex)

		for i := 0; i < len(slice); i++ {
			result += slice[shuffledSlice[i]]
		}
		if result == str {
			result = ""
			continue
		} else {
			break
		}

	}
	return result
}
