package main

import (
	"fmt"

	answers "github.com/skang0601/advent-of-code-2022/answers"
)

func main() {
	answerMap := map[string]func() error{
		"1": answers.One,
		"2": answers.Two,
		"3": answers.Three,
		"4": answers.Four,
		"5": answers.Five,
	}

	fmt.Println("Enter the advent day (1-25)")
	var input string
	fmt.Scanln(&input)

	if fn, ok := answerMap[input]; !ok {

	} else {
		err := fn()
		if err != nil {
			return
		}
	}
}
