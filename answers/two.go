package answers

import (
	"fmt"
	"strings"
)

func ScoreMatch() error {
	inputFile := "two.input"
	input, err := readInput(inputFile)

	if err != nil { return err }

	score := 0
	for _, i := range input {
		score += parseMatch(i)
	}

	fmt.Println(score)
	return nil
}


func parseMatch(match string) int {
	m := strings.Split(match, " ")

	fmt.Printf("%d-%d\n", ParseHand(m[0]), ParseHand(m[1]))
	//fmt.Printf("Hand %d + Match %d = %d\n",handScore(m[1]), matchScore(ParseHand(m[0]), ParseHand(m[1])), 0)

	//return handScore(m[1]) + matchScore(ParseHand(m[0]), ParseHand(m[1]))

	chosenCard, score := chooseCard(m[0], m[1])

	return chosenCard + score
}

func chooseCard(opponent, strategy string) (int, int) {
	switch strategy {
	case "Y":
		return ParseHand(opponent), 3
	case "Z":
		return ParseHand(opponent) + 1 % 3, 6
	default:
		if ParseHand(opponent) == 0 {
			return 2, 0
		}
		return ParseHand(opponent) - 1, 0

	}
}

func ParseHand(h string) int {
	switch h {
    case "A","X": // Rock
		return 0
	case "B","Y": // Paper
		return 1
	case "C","Z": // Scissors
		return 2
	default:
		fmt.Println("shouldn't happen")
		return -1

	}
}

func matchScore(opponent, hand int) int {
	/*
	   0 - Rock
	   1 - Paper
	   2 - Scissor
	  **/
	fmt.Println(-1 % 3)


	if hand == opponent {
		return 3
	}

	if (hand + 1) % 3 == opponent {
		return 0
	}


	if hand == 0 && opponent == 2 || (hand - 1) % 3 == opponent {
		return 6
	}

	return 0
}

func handScore(hand string) int {
	switch hand {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	default:
		return 0
	}
}
