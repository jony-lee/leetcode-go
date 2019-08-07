package main

import "fmt"

//【题目】
/*
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
// TODO
//【思路】
/*






 */

//【代码】
func longestValidParentheses(s string) int {
	var left, right, ret int
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == '(' {
			left++
		}
		if ch == ')' {
			right++
		}
		if left == right {
			if ret < left {
				ret = left
			}
		}
		if left < right {
			left = 0
			right = 0
		}
	}
	left = 0
	right = 0
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		if ch == '(' {
			left++
		}
		if ch == ')' {
			right++
		}
		if left == right {
			if ret < left {
				ret = left
			}
		}
		if left > right {
			left = 0
			right = 0
		}
	}

	return ret * 2
}

//【主函数】
func main() {
	a := ")(()())"
	fmt.Println(longestValidParentheses(a))
}

//【总结】：
/*







 */
