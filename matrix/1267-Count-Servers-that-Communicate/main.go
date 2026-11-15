package main

type Position struct {
	X int
	Y int
}

func countServers(grid [][]int) int {
	isolatedServer := 0
	total := 0
	xCheck := make([]int, len(grid))
	yCheck := make([]int, len(grid[0]))
	servers := make([]Position, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				total++
				servers = append(servers, Position{X: i, Y: j})
				xCheck[i]++
				yCheck[j]++
			}
		}
	}

	for _, v := range servers {
		if xCheck[v.X] == 1 && yCheck[v.Y] == 1 {
			isolatedServer++
		}
	}

	return total - isolatedServer
}
