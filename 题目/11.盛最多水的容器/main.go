package main

import "fmt"

//【题目】

/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
//【思路】
/*
方法一：暴力法
算法

在这种情况下，我们将简单地考虑每对可能出现的线段组合并找出这些情况之下的最大面积。

方法二：双指针法
算法

这种方法背后的思路在于，两线段之间形成的区域总是会受到其中较短那条长度的限制。此外，两线段距离越远，得到的面积就越大。

我们在由线段长度构成的数组中使用两个指针，一个放在开始，一个置于末尾。 此外，我们会使用变量 maxareamaxarea 来持续存储到目前为止所获得的最大面积。 在每一步中，我们会找出指针所指向的两条线段形成的区域，更新 maxareamaxarea，并将指向较短线段的指针向较长线段那端移动一步。

作者：LeetCode
链接：https://leetcode-cn.com/problems/two-sum/solution/sheng-zui-duo-shui-de-rong-qi-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

//【代码】
// 方法一 暴力解法
func maxArea1(height []int) int {
	m := 0
	a := 0
	b := 0
	for i := 0; i < len(height); i++ {
		a = height[i]
		for j := i + 1; j < len(height); j++ {
			b = height[j]
			if m < min(b, a)*(j-i) {
				m = min(b, a) * (j - i)
			}
		}
	}
	return m
}

func min(a, b int) int {
	if a > b {
		return min(b, a)
	}
	return a
}

// 方法二：双指针法
func maxArea(height []int) int {
	m := 0
	a := 0
	b := len(height) - 1
	for a < b {
		if m < min(height[a], height[b])*(b-a) {
			m = min(height[a], height[b]) * (b - a)
		}
		if height[a] < height[b] {
			a++
		} else {
			b--
		}
	}
	return m

}

//【主函数】
func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

//【总结】：
/*







 */
