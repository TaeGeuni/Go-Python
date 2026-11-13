package main

func countServers(grid [][]int) int {
	res := 0

	check := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		check[i] = make([]bool, len(grid[i]))
	}

	for i, _ := range grid {
		for j, w := range grid[i] {
			if w == 1 && !check[i][j] {
				check[i][j] = true
				queue
			}

			check[i][j] = true
		}
	}

	return res
}
