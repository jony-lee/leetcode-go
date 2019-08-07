package main

import "fmt"

//【题目】
/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:

输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
//
//【思路】
/*
1、找到最小值所在位置（二分查找）
2、以最小值划分两个有序子数组，判断target在哪个有序子数组中
3、二分查找target所在的位置



*/

//【代码】
func search(nums []int, target int) int {
	if len(nums) == 0 { //数组为空
		return -1
	}
	if len(nums) == 1 { //数组长度为1
		if nums[0] == target {
			return 0
		}
		return -1
	}
	index := searchMin(nums) //数组长度大于1，考虑旋转情况
	if index == 0 {          //无旋转情况
		return searchTarget(nums, target)
	}
	//有旋转
	if target == nums[0] {
		return 0
	} else if target > nums[0] { //如果目标大于第一个，则结果在前半段
		return searchTarget(nums[:index], target)
	} else { //如果目标小于第一个，则结果在后半段
		ret := searchTarget(nums[index:], target)
		if ret != -1 {
			ret += index //后半段添加数组偏移量
		}
		return ret
	}

}

func searchMin(nums []int) int { //找到最小值在nums中的位置
	left, right := 0, len(nums)-1
	if nums[left] < nums[right] {
		return 0
	}
	for left <= right {
		k := (left + right) / 2
		if nums[k] > nums[k+1] {
			return k + 1
		} else {
			if nums[k] < nums[left] {
				right = k - 1
			} else {
				left = k + 1
			}
		}
	}
	return 0
}

func searchTarget(nums []int, target int) int { //找到target在子数组中的位置
	left, right := 0, len(nums)-1
	for left <= right {
		k := (left + right) / 2
		if nums[k] > target {
			right = k - 1
		} else if nums[k] < target {
			left = k + 1
		} else {
			return k
		}
	}
	return -1
}

//【主函数】
func main() {
	a := []int{4, 5, 6, 7, 0, 1, 2}
	//fmt.Println(search(a,))
	target := 4
	fmt.Println(search(a, target))
}

//【总结】：
/*







 */
