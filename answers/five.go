package answers

import (
	"bufio"
	"fmt"
	"strings"
)

// --- Day 5: Supply Stacks ---

// The expedition can depart as soon as the final supplies have been unloaded from the ships. Supplies are stored in stacks of marked crates, but because the needed supplies are buried under many other crates, the crates need to be rearranged.

// The ship has a giant cargo crane capable of moving crates between stacks. To ensure none of the crates get crushed or fall over, the crane operator will rearrange them in a series of carefully-planned steps. After the crates are rearranged, the desired crates will be at the top of each stack.

// The Elves don't want to interrupt the crane operator during this delicate procedure, but they forgot to ask her which crate will end up where, and they want to be ready to unload them as soon as possible so they can embark.

// They do, however, have a drawing of the starting stacks of crates and the rearrangement procedure (your puzzle input). For example:

//     [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

// move 1 from 2 to 1
// move 3 from 1 to 3
// move 2 from 2 to 1
// move 1 from 1 to 2

// In this example, there are three stacks of crates. Stack 1 contains two crates: crate Z is on the bottom, and crate N is on top. Stack 2 contains three crates; from bottom to top, they are crates M, C, and D. Finally, stack 3 contains a single crate, P.

// Then, the rearrangement procedure is given. In each step of the procedure, a quantity of crates is moved from one stack to a different stack. In the first step of the above rearrangement procedure, one crate is moved from stack 2 to stack 1, resulting in this configuration:

// [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

// In the second step, three crates are moved from stack 1 to stack 3. Crates are moved one at a time, so the first crate to be moved (D) ends up below the second and third crates:

//         [Z]
//         [N]
//     [C] [D]
//     [M] [P]
//  1   2   3

// Then, both crates are moved from stack 2 to stack 1. Again, because crates are moved one at a time, crate C ends up below crate M:

//         [Z]
//         [N]
// [M]     [D]
// [C]     [P]
//  1   2   3

// Finally, one crate is moved from stack 1 to stack 2:

//         [Z]
//         [N]
//         [D]
// [C] [M] [P]
//  1   2   3

// The Elves just need to know which crate will end up on top of each stack; in this example, the top crates are C in stack 1, M in stack 2, and Z in stack 3, so you should combine these together and give the Elves the message CMZ.

// After the rearrangement procedure completes, what crate ends up on top of each stack?
//
//
// --- Part Two ---

// As you watch the crane operator expertly rearrange the crates, you notice the process isn't following your prediction.

// Some mud was covering the writing on the side of the crane, and you quickly wipe it away. The crane isn't a CrateMover 9000 - it's a CrateMover 9001.

// The CrateMover 9001 is notable for many new and exciting features: air conditioning, leather seats, an extra cup holder, and the ability to pick up and move multiple crates at once.

// Again considering the example above, the crates begin in the same configuration:

//     [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

// Moving a single crate from stack 2 to stack 1 behaves the same as before:

// [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

// However, the action of moving three crates from stack 1 to stack 3 means that those three moved crates stay in the same order, resulting in this new configuration:

//         [D]
//         [N]
//     [C] [Z]
//     [M] [P]
//  1   2   3

// Next, as both crates are moved from stack 2 to stack 1, they retain their order as well:

//         [D]
//         [N]
// [C]     [Z]
// [M]     [P]
//  1   2   3

// Finally, a single crate is still moved from stack 1 to stack 2, but now it's crate C that gets moved:

//         [D]
//         [N]
//         [Z]
// [M] [C] [P]
//  1   2   3

// In this example, the CrateMover 9001 has put the crates in a totally different order: MCD.

// Before the rearrangement process finishes, update your simulation so that the Elves know where they should stand to be ready to unload the final supplies. After the rearrangement procedure completes, what crate ends up on top of each stack?

func Five() error {
	inputs, err := readInput("five.input")
	output := ""
	if err != nil {
		return err
	}

	// Figure out where the initial stack state input ends
	// Parsing the input is the fun part of this problem lol
	// We know this ends when we encounter the first empty line in the input
	c := 0

	for i := 0; i < len(inputs); i++ {
		if inputs[i] == "" {
			c = i
			break
		}
	}

	// Skip the last line since it's just the labels
	stacks := parseInitialStacks(inputs[0 : c-1])

	// Start reading and performing in the stack instructions
	for i := c + 1; i < len(inputs); i++ {
		//parseStackInstruction(stacks, inputs[i])
		parseStackInstructionTwo(stacks, inputs[i])
	}

	// Construct the output
	for i := 0; i < len(stacks); i++ {
		output += stacks[i].Peek()
	}
	fmt.Println(output)

	return nil
}

func parseInitialStacks(inputs []string) map[int]*Stack[string] {
	stacks := map[int]*Stack[string]{}

	// Construct the stack(s) by reading backwards
	for i := len(inputs) - 1; i >= 0; i-- {
		// Parse out the stack state

		// Split based that separates stack entries knowing they're 3 character wide.
		// "   " is as valid entry for a stack as "[C]"

		// Can I do this with regex? I'm not actually sure...
		// fmt.Printf("%v\n", regexp.MustCompile(`\s?.{3}\s?`).Split(inputs[i], -1))
		// This is a pain in the ass, I'm just gonna use a custom scanner...

		// Otherwise just keep reading each entry manually.
		// Each entry is 3 runes separated by whitespaces so use a custom split function?
		scanner := bufio.NewScanner(strings.NewReader((inputs[i])))
		split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			// Each token is 3 characters long separated by a whitespace
			token = data[1:2]
			advance = 4

			// On the last entry, otherwise inputs get discarded
			if len(data) <= 3 {
				err = bufio.ErrFinalToken
			}
			return
		}
		scanner.Split(split)

		idx := 0
		for scanner.Scan() {
			if _, ok := stacks[idx]; !ok {
				stacks[idx] = NewStack[string]()
			}
			e := scanner.Text()
			if e != " " {
				stacks[idx].Push(e)
			}
			idx += 1
		}
	}
	return stacks
}

func parseStackInstruction(stacks map[int]*Stack[string], input string) {
	var numToMove, src, dst int

	if input == "" {
		return
	}

	_, err := fmt.Sscanf(input, "move %d from %d to %d", &numToMove, &src, &dst)

	if err != nil {
		panic(err)
	}

	for i := 0; i < numToMove; i++ {
		popped := stacks[src-1].Pop()

		if popped != "" {
			stacks[dst-1].Push(popped)
		}

	}
	return
}

// Just use an intermediatery stack to retain order
func parseStackInstructionTwo(stacks map[int]*Stack[string], input string) {
	var numToMove, src, dst int
	loading := []string{}

	if input == "" {
		return
	}

	_, err := fmt.Sscanf(input, "move %d from %d to %d", &numToMove, &src, &dst)

	if err != nil {
		panic(err)
	}

	for i := 0; i < numToMove; i++ {
		popped := stacks[src-1].Pop()

		if popped != "" {
			loading = append(loading, popped)
		}
	}

	for i := len(loading) - 1; i >= 0; i-- {
		stacks[dst-1].Push(loading[i])
	}
	return
}
