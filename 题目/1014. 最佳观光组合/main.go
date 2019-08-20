package main

//【题目】
/*
给定正整数数组 A，A[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的距离为 j - i。

一对景点（i < j）组成的观光组合的得分为（A[i] + A[j] + i - j）：景点的评分之和减去它们两者之间的距离。

返回一对观光景点能取得的最高分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-sightseeing-pair
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
//【思路】
/*

已知题目要求 res = A[i] + A[j] + i - j （i < j） 的最大值，

而对于输入中的每一个 A[j] 来说， 它的值 A[j] 和它的下标 j 都是固定的，

所以 A[j] - j 的值也是固定的。

因此，对于每个 A[j] 而言， 想要求 res 的最大值，也就是要求 A[i] + i （i < j） 的最大值，

所以不妨用一个变量 pre_max 记录当前元素 A[j] 之前的 A[i] + i 的最大值，

这样对于每个 A[j] 来说，都有 最大得分 = pre_max + A[j] - j,

再从所有 A[j] 的最大得分里挑出最大值返回即可。



作者：JiayangWu
链接：https://leetcode-cn.com/problems/best-sightseeing-pair/solution/python-jie-fa-by-jiayangwu/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。




*/

//【代码】
func maxScoreSightseeingPair(A []int) int {
	res := 0
	preMax := A[0] + 0
	for j := 1; j < len(A); j++ {
		// 找到 A[i] + A[j] + i - j （i < j）的最大值
		if res < preMax+A[j]-j {
			res = preMax + A[j] - j
		}
		// 找到 A[i] + i 的最大值
		if preMax < A[j]+j {
			preMax = A[j] + j
		}
	}
	return res
}

//【主函数】
func main() {

}

//【总结】：
/*







 */
