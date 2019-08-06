package main

import (
	"fmt"
)

/*
【题目】
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:

你能不将整数转为字符串来解决这个问题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

【思路】
//
简单题，反转一下数字即可。
因为如果是回文数的话，就不存在溢出，因此可以进行反转操作
如果发生溢出，则说明这个数不是回文数，返回false


*/

//【代码-数字实现-Jony】
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y, tmp := 0, x
	for tmp != 0 {
		y = y*10 + tmp%10
		tmp /= 10
	}
	return y == x
}

//【主函数】
func main() {
	a := 12211
	fmt.Println(isPalindrome(a))

}

/*
【总结】：
简单题





*/
