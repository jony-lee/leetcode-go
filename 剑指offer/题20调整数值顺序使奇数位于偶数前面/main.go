package main

import "fmt"

// 输入一个整数数组实现一个函数来调整该数组中,数字的顺序,使得所有奇数位于数组前半部分

// 使用两个指针,a在开头,b在末尾
// a向后移动,移动到偶数时停止
// 此时让b向前运动,移动到奇数时a和b的元素互换
// 然后再启动a向后移动,反复

func fun1(data []int) []int {
	a, b := 0, len(data)-1
	for a < b {
		if data[a]%2 == 1 {
			a++
		} else if data[b]%2 == 0 {
			b--
		} else {
			data[a], data[b] = data[b], data[a]
		}
	}
	return data
}

func main() {
	fmt.Print(fun1([]int{1, 2, 3, 4, 5, 6, 7}))

}
