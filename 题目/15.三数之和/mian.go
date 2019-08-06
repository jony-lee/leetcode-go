package main

import "sort"

//【题目】
/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
// TODO
//【思路】
/*

见注释

*/

//【代码】
func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 排序
	n := [][]int{}
	m := make(map[[3]int]int)          // 使用一个字典来去重
	for i := 1; i < len(nums)-1; i++ { // 从nums[1,len(s)} 中任选一个元素 i
		for v, u := 0, len(nums)-1; v < i && u > i; { // 将 最小的元素 v ,和最大的元素 u 与 i 相加 v + i + u
			a, b, c := nums[v], nums[u], nums[i]
			if nums[v]+nums[u]+nums[i] == 0 { // 如果等于 0 直接返回,并且 v++, u++,寻找新的合适组合
				if _, ok := m[[3]int{a, b, c}]; !ok { // 去重操作
					n = append(n, []int{nums[v], nums[u], nums[i]})
				}
				m[[3]int{a, b, c}] = 1
				v++
				u--
			} else if nums[v]+nums[u]+nums[i] < 0 { // 如果 v + i + u < 0 说明 v 太小了, 需要大一点的数, 则 v ++
				v++
			} else { // 同理 u--
				u--
			}
		}
	}
	return n
}

//【主函数】
func main() {

}

//【总结】：
/*







 */
