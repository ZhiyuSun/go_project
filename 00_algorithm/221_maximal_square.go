package main

import "fmt"

// 221. 最大正方形
// 2021.05.25 磕磕绊绊，终于写出来了
func maximalSquare(matrix [][]byte) int {
	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row)
	for i:=0; i<row; i++{
		dp[i] = make([]int, col)
	}
	maxV := 0
	for i:=0; i<row; i++ {
		for j:=0; j<col; j++ {
			if i == 0 || j == 0 {
				if matrix[i][j] == '1'{
					dp[i][j] = 1
				}else {
					dp[i][j] = 0
				}
			} else {
				if matrix[i][j] == '1' {
					dp[i][j] = minValue(dp[i-1][j-1], minValue(dp[i-1][j], dp[i][j-1])) + 1
				}
			}
			maxV = maxValue(dp[i][j], maxV)
		}
	}
	return maxV * maxV
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}


// 官方解法，简洁了不少
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxSide := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maximalSquare([][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}))
}
