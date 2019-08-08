package main

import "fmt"

//【题目】
/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-insert-position
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
//
//【思路】
/*
简单二分查找题，重点在于跳出边界条件的设定
当lo==hi时，lo肯定落在target上或target插入位置的左边
如果落在上面直接返回k
如果落在target插入位置的左边，不妨再跑一次，让lo++
因此，lo就是target最终插入时所在的位置。





*/

//【代码】
func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		k := (lo + hi) / 2
		if nums[k] > target {
			hi = k - 1
		} else if nums[k] < target {
			lo = k + 1
		} else {
			return k
		}
	}
	return lo
}

//【主函数】
func main() {
	a := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(a, target))
}

//【总结】：
/*







 */
