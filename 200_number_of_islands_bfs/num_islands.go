package numberOfIslands

const LAND = '1'
const SEA = '0'

func numIslands(grid [][]byte) int {
	islands := 0
	row := len(grid)
	if row < 1 {
		return islands
	}
	column := len(grid[0])
	if column < 1 {
		return islands
	}

	for r := range row {
		for c := range grid[r] {
			if grid[r][c] == SEA {
				continue
			}
			islands++
			floodFillBFS(grid, r, c)
		}
	}

	return islands
}

func floodFillBFS(grid [][]byte, row int, column int) {
	if row < 0 || column < 0 ||
		row >= len(grid) || column >= len(grid[row]) {
		return
	}
	queue := [][2]int{{row, column}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		r, c := current[0], current[1]
		points := [][2]int{{r - 1, c}, {r, c - 1}, {r + 1, c}, {r, c + 1}}
		for _, point := range points {
			pointRow := point[0]
			pointColumn := point[1]
			if pointRow < 0 || pointColumn < 0 ||
				pointRow >= len(grid) || pointColumn >= len(grid[pointRow]) {
				continue
			}
			if grid[pointRow][pointColumn] == LAND {
				grid[pointRow][pointColumn] = SEA
				queue = append(queue, [2]int{pointRow, pointColumn})
			}
		}
	}
}
