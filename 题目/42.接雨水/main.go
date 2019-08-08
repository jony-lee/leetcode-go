package main

//【题目】
/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
//
//【思路】
/*
1、找到最大值
2、从左向最大值处计算蓄水
3、从右向最大值处计算蓄水
4、返回蓄水量




*/

//【代码】
func trap(height []int) int {
	length := len(height)
	if length < 3 {
		return 0
	}

	//找到最大的柱子
	max := height[0]
	maxIndex := 0
	for i := 1; i < length; i++ {
		if height[i] > max {
			max = height[i]
			maxIndex = i
		}
	}
	volume := 0
	//从左边蓄水
	hi := height[0]
	for i := 1; i < maxIndex; i++ {
		if height[i] < hi {
			volume = volume + hi - height[i] //当前格子蓄水量
		} else {
			hi = height[i] //提高柱子
		}
	}
	//从右边蓄水
	hi = height[length-1]
	for i := length - 1; i > maxIndex; i-- {
		if height[i] < hi {
			volume = volume + hi - height[i] //当前格子蓄水量
		} else {
			hi = height[i] //提高柱子
		}
	}
	return volume
}

//【主函数】
func main() {

}

//【总结】：
/*







 */
