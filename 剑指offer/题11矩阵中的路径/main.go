package main

// 判断一个矩阵中是否存在一条包含某字符串所有字符的路径,可以从任意一单元开始
// 每一步只能上下左右移动一个单元,每个单元只能被访问一次

func hasPath(matrix string, rows, cols int, s string) bool {
	visited := make([]int, rows*cols)
	index := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if hasCode(matrix, rows, cols, row, col, s, &index, visited) {
				return true
			}
		}
	}
	return false
}

func hasCode(matrix string, rs int, cs int, r int, c int, str string, index *int, visited []int) bool {
	if *index == len(matrix) {
		return true
	}
	hasPath := false
	if r >= 0 && r < rs && c >= 0 && c < cs && matrix[r*cs+c] == str[*index] && visited[r*cs+c] == 0 {
		*index++
		visited[r*cs+c] = 1
		hasPath = hasCode(matrix, rs, cs, r, c-1, str, index, visited) ||
			hasCode(matrix, rs, cs, r, c+1, str, index, visited) ||
			hasCode(matrix, rs, cs, r-1, c, str, index, visited) ||
			hasCode(matrix, rs, cs, r+1, c, str, index, visited)
		if !hasPath {
			*index--
			visited[r*cs+c] = 0
		}
	}
	return hasPath
}
