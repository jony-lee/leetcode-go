package main

import (
	"fmt"
	"sort"
)

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

//【思路】
/*
先考虑对数组进行排序，然后选择固定一个数，去寻找另外两个匹配的数
题目的关键点在于去重，代码1是选择了固定中心点然后向两端寻求匹配数


其他见注释
*/

//【代码】
func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 排序
	var n [][]int
	m := make(map[[3]int]int)          // 使用一个字典来去重
	for i := 1; i < len(nums)-1; i++ { // 从nums[1,len(s)} 中任选一个元素 i
		for v, u := 0, len(nums)-1; v < i && u > i; { // 将 最小的元素 v ,和最大的元素 u 与 i 相加 v + i + u
			a, b, c := nums[v], nums[u], nums[i]
			if a+b+c == 0 { // 如果等于 0 直接返回,并且 v++, u++,寻找新的合适组合
				if _, ok := m[[3]int{a, b, c}]; !ok { // 去重操作
					n = append(n, []int{a + b + c})
				}
				m[[3]int{a, b, c}] = 1
				v++
				u--
			} else if a+b+c < 0 { // 如果 v + i + u < 0 说明 v 太小了, 需要大一点的数, 则 v ++
				v++
			} else { // 同理 u--
				u--
			}
		}
	}
	return n
}

//【代码1】
func threeSum1(nums []int) [][]int {
	sort.Ints(nums) //原址排序
	var ret [][]int
	m := make(map[[3]int]int) //要点：字典是要求key必须能支持== != 运算符的数据类型
	for i := 0; i < len(nums)-1 && nums[i] <= 0; i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			a, b, c := nums[i], nums[j], nums[k]
			if a+b+c == 0 {
				if _, ok := m[[3]int{a, b, c}]; !ok {
					ret = append(ret, []int{a, b, c})
				}
				m[[3]int{a, b, c}] = 1
				j++
				k--
			} else if a+b+c < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return ret
}

//【代码2】
func threeSum2(nums []int) [][]int {
	sort.Ints(nums) //排序
	var result [][]int
	for i := 0; i < len(nums)-1 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] { //第一次循环不执行，后面去重复
			continue
		}
		j := i + 1 //双指针
		k := len(nums) - 1
		target := 0 - nums[i]
		for j < k {
			sum := nums[j] + nums[k]
			if sum == target { //如果相等，找到对应的数
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
				k--
				for j < k && nums[k+1] == nums[k] {
					k--
				}
			} else if sum > target { //如果大了就找一个小的数
				k--
				for j < k && nums[k+1] == nums[k] {
					k--
				}
			} else if sum < target { //如果小了就找一个大的数
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
			}
		}
	}
	return result
}

//【主函数】
func main() {
	a := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(a))
}

//【总结】：
/*
利用双指针进行求解

代码2和代码3选择了固定左边点，然后从两端向中心寻求匹配数
代码1和代码2选择的去重方式是通过字典进行去重，字典去重要求key必须是能支持== != 运算符的数据类型（这一块卡了半天）
代码3直接是通过for循环判断去实现的
上述三段代码各有优劣势
代码1：
优势：利用字典去重，理解起来较容易
劣势：中心向两边扩散，没有充分利用排序的优势
劣势：字典额外占用内存，实际上是利用空间换时间
代码2：
优势：固定左端，两边向中心扩散，利用了左端大于0则跳出可对代码进行部分优化
优势：利用字典去重，理解起来较容易
劣势：字典额外占用内存，实际上是利用空间换时间
代码3：
优势：固定左端，两边向中心扩散，利用了左端大于0则跳出可对代码进行部分优化
优势：手动处理重复问题，代码更高效
劣势：对边界问题处理需小心，理解难度相对大
*/
