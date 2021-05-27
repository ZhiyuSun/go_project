package main

import "fmt"

// 200. 岛屿数量

// 2021.05.22 我自己的方法
type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func numIslands(grid [][]byte) int {
	row, col, res := len(grid), len(grid[0]), 0

	var dfs func(i, j int)
	dfs = func(i, j int){
		grid[i][j] = '0'
		for _, d := range directions{
			ni := i + d.x
			nj := j + d.y
			if ni >= 0 && ni < row && nj >= 0 && nj < col && grid[ni][nj] == '1'{
				dfs(ni, nj)
			}
		}
	}

	for i:=0; i<row; i++ {
		for j:=0; j<col; j++ {
			if grid[i][j] == '1'{
				res ++
				dfs(i, j)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1' }}))
}
