package main

import "fmt"

// 在长度为n+1的数组里的所有数字都在1~n范围内,找出任意一个重复的数字,不能修改数组
// 失败返回-1

func getDuplication(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] < 1 || nums[i] > len(nums)-1 {
			return -1
		}
	}
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = 0
		} else {
			return nums[i]
		}
	}
	return -1
}

func main() {
	fmt.Println(getDuplication([]int{1, 2, 2, 2}))
}
