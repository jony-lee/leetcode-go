package main

import "sort"

//【题目】
/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
// TODO
//【思路】
/*


见三数之和


*/

//【代码】
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)                    //排序
	n := nums[1] + nums[0] + nums[2]   // 初始化三数之和
	for i := 1; i < len(nums)-1; i++ { // 从nums[1,len(s)} 中任选一个元素 i
		for a, b := 0, len(nums)-1; a < i && b > i; {
			i, j, k := nums[i], nums[a], nums[b]
			if i+j+k == target {
				return target
			} else if i+j+k > target {
				b--
				if abs(i+j+k-target) < abs(n-target) {
					n = i + j + k
				}
			} else {
				a++
				if abs(i+j+k-target) < abs(n-target) {
					n = i + j + k
				}
			}
		}
	}
	return n
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//【主函数】
func main() {

}

//【总结】：
/*







 */
