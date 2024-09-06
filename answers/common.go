package answers

import (
	"bufio"
	"os"
)

/*
 * Common utility functions I'll probably fucking need because go has garbage libraries
 */

const (
	inputDir = "./inputs/"
)

func readInput(fileName string) ([]string, error) {
	f, err := os.Open(inputDir + fileName)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	output := make([]string, 0)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return output, nil
}
