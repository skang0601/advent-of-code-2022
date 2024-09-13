package answers

import (
	"bufio"
	"os"
)

/*
 * Common utility functions I'll probably fucking need because go has garbage std libraries
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

// Good excuse as any to play around with golang Generics and make a Stack impl I guess ¯\_(ツ)_/¯
type Stack[T any] struct {
	entries []T
}

// LOL fuck me, they don't allow for type parameters on methods
//func (s Stack) Insert[T string|int](i T) {

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		entries: nil,
	}
}

func (s *Stack[T]) Push(i T) {
	s.entries = append(s.entries, i)
}

func (s *Stack[T]) Pop() (t T) {
	if len(s.entries) == 0 {
		return t
	}

	v := s.entries[len(s.entries)-1]
	s.entries = s.entries[:len(s.entries)-1]

	return v
}

func (s *Stack[T]) Peek() (t T) {
	if len(s.entries) == 0 {
		return t
	}

	return s.entries[len(s.entries)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.entries)
}

func (s *Stack[T]) Copy() *Stack[T] {
	entries := make([]T, len(s.entries))
	copy(entries, s.entries)
	return &Stack[T]{
		entries: entries,
	}
}
