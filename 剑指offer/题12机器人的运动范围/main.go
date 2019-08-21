package main

// 地上有一个m行n列的方格,一个机器人从坐标(0,0),的格子开始移动,每次可以上下左右移动一次,
// 但不能进入行坐标和列坐标的数位之和大于k的格子,问机器人可以达到多少个格子?

func movingCount(threhold int, rows int, cols int) int {
	if threhold < 0 || rows <= 0 || cols <= 0 {
		return 0
	}
	visited := make([]int, rows*cols)
	count := movingCountCore(threhold, rows, cols, 0, 0, visited)
	return count
}

func movingCountCore(threhold int, rows int, cols int, row int, col int, visited []int) int {
	count := 0
	if check(threhold, rows, cols, row, col, visited) {
		visited[row*cols+col] = 1
		count = 1 + movingCountCore(threhold, rows, cols, row-1, col, visited) +
			movingCountCore(threhold, rows, cols, row, col-1, visited) +
			movingCountCore(threhold, rows, cols, row+1, col, visited) +
			movingCountCore(threhold, rows, cols, row, col+1, visited)
	}
	return count
}

// 判断机器人是否能进入(row, col)方格
func check(threhold int, rows int, cols int, row int, col int, visited []int) bool {
	if row >= 0 && row < rows && col <= 0 && col < cols &&
		getDS(row)+getDS(col) <= threhold && visited[row*cols+col] == 0 {
		return true
	}
	return false
}

// 一个数字的位数之和
func getDS(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
