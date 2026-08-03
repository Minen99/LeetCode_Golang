package numberOfIslands

const LAND = '1'
const SEA = '0'

func numIslands(grid [][]byte) int {
	numOfIslands := 0
	row := len(grid)
	if row < 1 {
		return numOfIslands
	}
	column := len(grid[0])
	if column < 1 {
		return numOfIslands
	}

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == SEA {
				continue
			}
			numOfIslands++
			oneIsland(&grid, r, c)
		}
	}
	return numOfIslands
}

func oneIsland(grid *[][]byte, row int, column int) {
	// out of range
	if row < 0 || column < 0 ||
		row >= len(*grid) || column >= len((*grid)[row]) {
		return
	}
	// is already searched or is sea
	if (*grid)[row][column] == SEA {
		return
	}
	(*grid)[row][column] = SEA
	// each faces
	oneIsland(grid, row-1, column)
	oneIsland(grid, row, column-1)
	oneIsland(grid, row+1, column)
	oneIsland(grid, row, column+1)
}
