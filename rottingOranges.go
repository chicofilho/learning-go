/*
Leet code problem: https://leetcode.com/problems/rotting-oranges/
In a given grid, each cell can have one of three values:

the value 0 representing an empty cell;
the value 1 representing a fresh orange;
the value 2 representing a rotten orange.
Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange.
If this is impossible, return -1 instead.
*/
package learninggo

// helper struct to make the stack implementation and find the min in a BFS
type posMin struct {
	i   int
	j   int
	min int
}

//search function to find the minimum minutes needed
func minMinutes(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}

	stack := getFirstRottenPositions(grid)
	minutes := 0
	for len(stack) > 0 {
		el := stack[0]
		stack = stack[1:]

		if el.min > minutes {
			minutes = el.min
		}
		for _, val := range getPositions(el.i, el.j) {
			ni, nj := val[0], val[1]
			if posInBoundaries(ni, nj, grid) && isFresh(ni, nj, grid) {
				grid[ni][nj] = 2
				stack = append(stack, posMin{ni, nj, el.min + 1})
			}
		}
	}

	if hasFreshOrange(grid) {
		return -1
	} else {
		return minutes
	}
}

// helper to get next positions
func getPositions(i int, j int) [][]int {
	return [][]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
}

// helper check to see if a position is inside boundaries
func posInBoundaries(ni int, nj int, grid [][]int) bool {
	return ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[ni])
}

// helper query for a fresh orange
func isFresh(ni int, nj int, grid [][]int) bool {
	return grid[ni][nj] == 1
}

// one pass through the grid to fetch the first positions
// all positions that have 2
func getFirstRottenPositions(grid [][]int) []posMin {
	res := []posMin{}
	for i, line := range grid {
		for j, val := range line {
			if val == 2 {
				res = append(res, posMin{i, j, 0})
			}
		}
	}
	return res
}

// one pass through the grid to check if any fresh orange is present
func hasFreshOrange(grid [][]int) bool {
	for _, line := range grid {
		for _, val := range line {
			if val == 1 {
				return true
			}
		}
	}
	return false
}
