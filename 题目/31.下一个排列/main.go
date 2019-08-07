package main

import (
	"fmt"
)

//【题目】
/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

//【思路】
/*
https://pic.leetcode-cn.com/1df4ae7eb275ba4ab944521f99c84d782d17df804d5c15e249881bafcf106173-file_1555696082944
图解思路很清楚了





*/

//【代码】
func nextPermutation(nums []int) {
	length := len(nums)
	if length == 1 {
		return
	}
	for i := length - 1; i >= 0; i-- {
		if i == 0 { //遍历到i==0的话说明是降序，直接反转后返回即可
			reverse(nums[:])
			return
		}
		if nums[i-1] >= nums[i] {
			continue
		}
		for j := i; j < length; j++ { //寻找只比i大一点点的数进行交换
			if i == length-1 { //说明最后一个数比倒数第二个数小，交换这两个数即可结束
				reverse(nums[length-2 : length])
				return
			}
			if nums[j] <= nums[i-1] { //从i右边开始寻找第一个比i小（或等）的数，那么j左边那个数即满足交换条件
				nums[i-1], nums[j-1] = nums[j-1], nums[i-1]
				reverse(nums[i:length])
				return
			}
			if j == length-1 { //如果一直找到头都没有比i小（或等的数），那么直接让最后一个数和交换数交换
				nums[i-1], nums[j] = nums[j], nums[i-1]
				reverse(nums[i:length])
				return
			}
		}

	}

}
func reverse(m []int) {
	i, j := 0, len(m)-1
	for i < j {
		m[i], m[j] = m[j], m[i]
		i++
		j--
	}
}

//【主函数】
func main() {
	a := []int{1, 3, 2}
	nextPermutation(a)
	//reverse(a)
	fmt.Println(a)

}

//【总结】：
/*







 */
