package main

import "fmt"

// 找出数组中重复的数字
// 在一个长度为n的数组里,所有数字都在0~n-1的范围内,请找出数组中任意一个重复的数组

// 思路
// 遍历数组,当扫描到下标为i的数字m时,首先比较这个数字是不是等于i,如果是继续迭代
// 如果不是, 拿他和下标为m的数字比较,如果相等就找到了一个数字,返回m,如果不相等,就把
// 下标为i和m的数字交换反复,
// 查找失败或者数组不符合规则时返回-1

func duplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] > len(nums)-1 {
			return -1
		}
	}
	for i := 0; i < len(nums); i++ {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			} else {
				a := nums[i]
				nums[i] = nums[a]
				nums[a] = a
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(duplicate([]int{0, 1, 2, 2}))
}
