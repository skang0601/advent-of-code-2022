package answers

import (
	"bufio"
	"cmp"
	"os"
)

/*
 * Common utility functions I'll probably need I guess
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

// UGH they don't allow for type parameters on methods
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

type Node struct {
	children map[string]*Node
	// Useful for Problem 7 when constructing our tree
	parent *Node
	// Metadata for Problem 7
	// Note: This is meant to represent the size in bytes in a file system.
	// Should be set to 0 if it's a directory
	size int
	name string
}

func NewNode(size int, name string) *Node {
	root := Node{
		children: make(map[string]*Node, 0),
		parent:   nil,
		size:     size,
		name:     name,
	}
	return &root
}

func (n *Node) Size() int {
	return n.size
}

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) Add(m *Node) {
	n.children[m.name] = m
	m.parent = n

	// This may not enough?
	// Do we want to update the total size of the tree on every insertion or just calculate when we need it?
	// I think it's better to do this calculation with the WalkDir.
	// n.size += m.size

}

func (n *Node) IsDir() bool {
	return len(n.children) != 0
}

// This skips the root dir
// This isn't how this is implemented in the std library, might be fun to peek at that impl later.
func (n *Node) WalkDir(fn func(m *Node)) {
	for _, v := range n.children {
		// Skip non-directory
		if len(v.children) == 0 {
			continue
		}
		// Recurse into each child
		v.WalkDir(fn)

		// Run the passed in function on each children
		fn(v)
	}
}

func (n *Node) Cd(path string) *Node {
	curr := n

	switch path {
	case "..":
		return curr.parent
	case "/":
		for curr.parent != nil {
			curr = curr.parent
		}
	default:
		if v, ok := curr.children[path]; ok {
			curr = v
		} else {
			// Should return an error really
			curr = nil
		}
	}

	return curr
}

func max[T cmp.Ordered](i, j T) T {
	if i > j {
		return i
	}

	return j
}
