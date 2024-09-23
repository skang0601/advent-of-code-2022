package answers

import (
	"fmt"
)

// --- Day 8: Treetop Tree House ---

// The expedition comes across a peculiar patch of tall trees all planted carefully in a grid. The Elves explain that a previous expedition planted these trees as a reforestation effort. Now, they're curious if this would be a good location for a tree house.

// First, determine whether there is enough tree cover here to keep a tree house hidden. To do this, you need to count the number of trees that are visible from outside the grid when looking directly along a row or column.

// The Elves have already launched a quadcopter to generate a map with the height of each tree (your puzzle input). For example:

// 30373
// 25512
// 65332
// 33549
// 35390

// Each tree is represented as a single digit whose value is its height, where 0 is the shortest and 9 is the tallest.

// A tree is visible if all of the other trees between it and an edge of the grid are shorter than it. Only consider trees in the same row or column; that is, only look up, down, left, or right from any given tree.

// All of the trees around the edge of the grid are visible - since they are already on the edge, there are no trees to block the view. In this example, that only leaves the interior nine trees to consider:

//     The top-left 5 is visible from the left and top. (It isn't visible from the right or bottom since other trees of height 5 are in the way.)
//     The top-middle 5 is visible from the top and right.
//     The top-right 1 is not visible from any direction; for it to be visible, there would need to only be trees of height 0 between it and an edge.
//     The left-middle 5 is visible, but only from the right.
//     The center 3 is not visible from any direction; for it to be visible, there would need to be only trees of at most height 2 between it and an edge.
//     The right-middle 3 is visible from the right.
//     In the bottom row, the middle 5 is visible, but the 3 and 4 are not.

// With 16 trees visible on the edge and another 5 visible in the interior, a total of 21 trees are visible in this arrangement.

// Consider your map; how many trees are visible from outside the grid?

// --- Part Two ---

// Content with the amount of tree cover available, the Elves just need to know the best spot to build their tree house: they would like to be able to see a lot of trees.

// To measure the viewing distance from a given tree, look up, down, left, and right from that tree; stop if you reach an edge or at the first tree that is the same height or taller than the tree under consideration. (If a tree is right on the edge, at least one of its viewing distances will be zero.)

// The Elves don't care about distant trees taller than those found by the rules above; the proposed tree house has large eaves to keep it dry, so they wouldn't be able to see higher than the tree house anyway.

// In the example above, consider the middle 5 in the second row:

// 30373
// 25512
// 65332
// 33549
// 35390

//     Looking up, its view is not blocked; it can see 1 tree (of height 3).
//     Looking left, its view is blocked immediately; it can see only 1 tree (of height 5, right next to it).
//     Looking right, its view is not blocked; it can see 2 trees.
//     Looking down, its view is blocked eventually; it can see 2 trees (one of height 3, then the tree of height 5 that blocks its view).

// A tree's scenic score is found by multiplying together its viewing distance in each of the four directions. For this tree, this is 4 (found by multiplying 1 * 1 * 2 * 2).

// However, you can do even better: consider the tree of height 5 in the middle of the fourth row:

// 30373
// 25512
// 65332
// 33549
// 35390

//     Looking up, its view is blocked at 2 trees (by another tree with a height of 5).
//     Looking left, its view is not blocked; it can see 2 trees.
//     Looking down, its view is also not blocked; it can see 1 tree.
//     Looking right, its view is blocked at 2 trees (by a massive tree of height 9).

// This tree's scenic score is 8 (2 * 2 * 1 * 2); this is the ideal spot for the tree house.

// Consider each tree on your map. What is the highest scenic score possible for any tree?

func Eight() error {
	input, err := readInput("eight.input")
	if err != nil {
		return err
	}
	total, maxScenicScore := 0, 0

	// We need to keep track and build up the max height table.
	// (i, j) is visible from an edge when...
	// 1. A border cell (h, k) is visible from and edge AND the value[i][j] > value[h][k]
	// 2. All cells to one of the edges from (i,j) is less than the value[i][j]
	// So is visible(i,j) = (memo[i-1][j] && input[i][j] > input[i-1][j]) || (memo[i][j-1] && input[i][j] > input[i][j-1])...
	// OR
	// values[0,...,i-1, i+1,...][j] < values[i][j] OR values[i][0,...,j-1,j+1,...] < values[i][j]
	// The above two actually just means that the value facing towards the (i,j) cell are all strictly less than (i,j)
	// 1. Are all values in memo[0][j]...memo[i-1][j] all less than memo[i][j]?
	// 2. Are all values in memo[i][0]...memo[i][j-1] all less than memo[i][j]?
	// 3. Are all values in memo[n][j]...memo[i+1][j] all less than memo[i][j]?
	// 4. Are all values in memo[i][n]...memo[i][j+1] all less than memo[i][j]?
	//
	// Ideas:
	// Going one direction for now
	// What if we store in memo[i][j] the largest encountered number we've seen and use that to determine if the tree is visible*
	// This means we'd need 4 copies right? 0~>i, 0~>j, n~>i, n~>j
	// But we only need two n*m loops ~> This can be collapsed down
	// Q:How do I detect whether input[i][j] is the biggest value in memo[i][j] or if there's a dupe?
	// A: memo[i][j] should contain the maximal value excluding the (i, j)th element.
	maxHeight := make([][]struct{ i, j, ni, nj byte }, len(input))

	for i := range maxHeight {
		maxHeight[i] = make([]struct{ i, j, ni, nj byte }, len(input))
	}

	for i := 0; i < len(maxHeight); i++ {
		for j := 0; j < len(maxHeight); j++ {
			if i == 0 {
				maxHeight[i][j].i = byte(0)
				maxHeight[len(maxHeight)-1][j].ni = byte(0)
			} else {
				maxHeight[i][j].i = max(input[i-1][j], maxHeight[i-1][j].i)
				maxHeight[len(maxHeight)-1-i][j].ni = max(input[len(maxHeight)-i][j], maxHeight[len(maxHeight)-i][j].ni)
			}

			if j == 0 {
				maxHeight[i][j].j = byte(0)
				maxHeight[i][len(maxHeight)-1].j = byte(0)
			} else {
				maxHeight[i][j].j = max(input[i][j-1], maxHeight[i][j-1].j)
				maxHeight[i][len(maxHeight)-1-j].nj = max(input[i][len(maxHeight)-j], maxHeight[i][len(maxHeight)-j].nj)
			}
		}
	}

	for i := 0; i < len(maxHeight); i++ {
		for j := 0; j < len(maxHeight); j++ {
			visible := maxHeight[i][j].i < input[i][j] || maxHeight[i][j].j < input[i][j] || maxHeight[i][j].ni < input[i][j] || maxHeight[i][j].nj < input[i][j]
			if visible {
				total += 1
			}
		}
	}

	/*
		fmt.Println("Max Height Map")
		for i := 0; i < len(maxHeight); i++ {
			for j := 0; j < len(maxHeight); j++ {
				fmt.Printf("{i: %s j: %s ni: %s nj: %s} ", string(maxHeight[i][j].i), string(maxHeight[i][j].j), string(maxHeight[i][j].ni), string(maxHeight[i][j].nj))
			}
			fmt.Print("\n")
		}
	*/

	fmt.Printf("Part 1: %d\n", total)

	// Part 2
	// For each inner entry calculate the scenic score.
	// Can I do better than this brute force method?
	// This is what? O((n * m)^2)
	// Could I pre-process the input like before?
	// If I store the l, r, u, d distance for each element is that reusable?
	//
	// Let's think about calculating the left distance on a 1d array..
	// if input[i] > input[i-1] then:
	//   1. input[i].l is at least input[i-1].l + 1
	//   q. Could it be more?
	//   a. Yes ~> [ 0 1 4 3 6 ]
	//      input[3].l = 1 but input[4].l = 4
	//   Does the maximum map help us? We know the biggest value from 0...i-1
	//   if the max < input[i][j] ~> input[i][j].l is i
	//   else ~> gets hairy as there might be other elements that's smaller than the max but bigger than input[i][j]
	//   ex: [ 0 5 4 3 ] ~> input[2] l_max is 5 but our view stops at 4
	//
	//   q:  we need to look at the maximum in both direction?
	//   a: No I don't think that helps in evaluating anything...
	//
	//   q: Should we store more state in the pre-processing?
	//   We could store multiple maximas maybe per height? The heights are well bounded....
	//   m[i][j] = []int

	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input)-1; j++ {
			l, r, u, d, h := 1, 1, 1, 1, input[i][j]
			for (j-l > 0 && input[i][j-l] < h) ||
				(j+r < len(input)-1 && input[i][j+r] < h) ||
				(i-u > 0 && input[i-u][j] < h) ||
				(i+d < len(input)-1 && input[i+d][j] < h) {

				if j-l > 0 && input[i][j-l] < h {
					//fmt.Printf("Left: %s is smaller than %s\n ", string(input[i][j-l]), string(h))
					l += 1
				}
				if j+r < len(input)-1 && input[i][j+r] < h {
					//fmt.Printf("Right: %s is smaller than %s\n ", string(input[i][j+r]), string(h))
					r += 1
				}
				if i-u > 0 && input[i-u][j] < h {
					//fmt.Printf("Up: %s is smaller than %s\n ", string(input[i-u][j]), string(h))
					u += 1
				}
				if i+d < len(input)-1 && input[i+d][j] < h {
					//fmt.Printf("Down: %s is smaller than %s\n ", string(input[i+d][j+d]), string(h))
					d += 1
				}
			}
			maxScenicScore = max(maxScenicScore, l*r*u*d)
		}
	}
	fmt.Printf("Part 2: %d\n", maxScenicScore)
	return nil
}
