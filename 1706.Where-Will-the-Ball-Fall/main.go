package main

import "fmt"

func findBall(grid [][]int) []int {
	n := len(grid[0])
	ans := make([]int, n)
	for i := range ans {
		col := i
		for _, row := range grid {
			dir := row[col]
			col += dir
			if col < 0 || col == n || row[col] != dir {
				col = -1
				break
			}
		}
		ans[i] = col
	}
	return ans

}

func main() {
	grid := [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}
	fmt.Printf("%v", findBall(grid))

}
