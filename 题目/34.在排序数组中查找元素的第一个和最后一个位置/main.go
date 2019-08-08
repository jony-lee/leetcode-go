package main

import "fmt"

//【题目】
/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
//
//【思路】
/*
两次二分查找，分别查找它的左边界和右边界，
重点在于返回条件





*/

//【代码】
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if target == nums[0] {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}

	lo, hi := 0, len(nums)-1
	var left, right int
	//查找左边界
	for lo < hi || nums[hi] != target {
		k := (lo + hi) / 2 //取中心或中心偏左的数
		if nums[k] < target {
			lo = k + 1
		} else if nums[k] > target {
			hi = k - 1
		} else {
			hi = k
		}
		if lo > hi {
			return []int{-1, -1}
		}
	}
	left = hi
	lo, hi = 0, len(nums)-1
	//查找右边界
	for lo < hi || nums[lo] != target {
		k := (lo+hi)/2 + 1 //取中心或中心偏右的数
		if nums[k] < target {
			lo = k + 1
		} else if nums[k] > target {
			hi = k - 1
		} else {
			lo = k
		}
		if lo > hi {
			return []int{-1, -1}
		}
	}
	right = lo
	return []int{left, right}
}

//【主函数】
func main() {
	a := []int{1, 2, 4, 4, 5, 6}
	fmt.Println(searchRange(a, 4))
}

//【总结】：
/*







 */
