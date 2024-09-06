package answers

import (
	"fmt"

	"golang.org/x/tools/go/analysis/passes/slog"
)

// --- Day 3: Rucksack Reorganization ---

// One Elf has the important job of loading all of the rucksacks with supplies for the jungle journey. Unfortunately, that Elf didn't quite follow the packing instructions, and so a few items now need to be rearranged.

// Each rucksack has two large compartments. All items of a given type are meant to go into exactly one of the two compartments. The Elf that did the packing failed to follow this rule for exactly one item type per rucksack.

// The Elves have made a list of all of the items currently in each rucksack (your puzzle input), but they need your help finding the errors. Every item type is identified by a single lowercase or uppercase letter (that is, a and A refer to different types of items).

// The list of items for each rucksack is given as characters all on a single line. A given rucksack always has the same number of items in each of its two compartments, so the first half of the characters represent items in the first compartment, while the second half of the characters represent items in the second compartment.

// For example, suppose you have the following list of contents from six rucksacks:

// vJrwpWtwJgWrhcsFMMfFFhFp
// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
// PmmdzqPrVvPwwTWBwg
// wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
// ttgJtRGJQctTZtZT
// CrZsJsPPZsGzwwsLwLmpwMDw

//     The first rucksack contains the items vJrwpWtwJgWrhcsFMMfFFhFp, which means its first compartment contains the items vJrwpWtwJgWr, while the second compartment contains the items hcsFMMfFFhFp. The only item type that appears in both compartments is lowercase p.
//     The second rucksack's compartments contain jqHRNqRjqzjGDLGL and rsFMfFZSrLrFZsSL. The only item type that appears in both compartments is uppercase L.
//     The third rucksack's compartments contain PmmdzqPrV and vPwwTWBwg; the only common item type is uppercase P.
//     The fourth rucksack's compartments only share item type v.
//     The fifth rucksack's compartments only share item type t.
//     The sixth rucksack's compartments only share item type s.

// To help prioritize item rearrangement, every item type can be converted to a priority:

//     Lowercase item types a through z have priorities 1 through 26.
//     Uppercase item types A through Z have priorities 27 through 52.

// In the above example, the priority of the item type that appears in both compartments of each rucksack is 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s); the sum of these is 157.

// Find the item type that appears in both compartments of each rucksack. What is the sum of the priorities of those item types?

func Three() error {
	inputFile := "three.input"
	inputs, err := readInput(inputFile)
	if err != nil {
		return err
	}

	sum := 0

	for _, i := range inputs {
		//fmt.Printf("Processing: %s\n", i)
		sum += toPriority(findCommonItem([]byte(i[0:len(i)/2]), []byte(i[len(i)/2:len(i)])))
	}

	fmt.Printf("Part 1: %d\n", sum)

	return nil
}

func findCommonItem(s1 []byte, s2 []byte) int {
	fmt.Printf("Comparing rucksack %s against %s\n", string(s1), string(s2))
	var found int
	seen := map[byte]bool{}

	for _, i := range s1 {
		seen[i] = true
	}

	for _, i := range s2 {
		if _, ok := seen[i]; ok {
			fmt.Printf("Found common element %s\n", string(i))
			found = int(i)
		}
	}

	return found
}

// Given the ascii character convert to a priority based on the rule (a-z) => (1-26) & (A-Z) => (27-52)
func toPriority(c int) int {
	// Just some ascii math
	if c >= 65 && c <= 90 {
		return (c - 65) + 27
	} else {
		return (c - 97) + 27
	}
}
