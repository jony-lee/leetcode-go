package main

import "sort"

//【题目】
/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
// TODO
//【思路】
/*






 */

//【代码】
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}
	m := map[[4]int]int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k, l := j+1, len(nums)-1; k < l; {
				a, b, c, d := nums[i], nums[j], nums[k], nums[l]
				if a+b+c+d == target {
					if _, ok := m[[4]int{a, b, c, d}]; !ok {
						ret = append(ret, []int{a, b, c, d})
					}
					m[[4]int{a, b, c, d}] = 1
					k++
					l--
				} else if a+b+c+d < target {
					k++
				} else {
					l--
				}
			}

		}
	}
	return ret
}

//【主函数】
func main() {

}

//【总结】：
/*







 */
