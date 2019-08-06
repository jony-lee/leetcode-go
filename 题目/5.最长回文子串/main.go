package main

import "fmt"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

解题思路：
//TODO Manacher算法用于解该题最好，但是没看懂，之后再看吧
看到最长回文子串，我第一想法就是先看看之前的无重复子串的思路能不能用上
同样的，
时间效率O()
空间效率O()
总结：
*/

//思路1-------------------------------------------

func longestPalindrome(s string) string {
	start, end, length, max := 0, 0, 0, 0
	m := make(map[int32]int)
	for i, v := range s {

		if old, ok := m[v]; ok == true && start < old+1 {
			start = old + 1
		}
		end = i
		m[v] = i
		length = end - start + 1
		if max < length {
			max = length
		}
	}
	return max
}
func longestPalindrome1(s string) string { // 中心扩展法
	if len(s) < 1 {
		return ""
	}
	st, e := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)   // 奇数长回文的中心
		len2 := expandAroundCenter(s, i, i+1) // 偶数长回文的中心
		lenm := max(len1, len2)
		if lenm > e-st {
			st = i - (lenm-1)/2
			e = i + lenm/2
		}
	}
	return s[st : e+1]
}

func expandAroundCenter(s string, left int, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

//他人优秀代码----------------------------------

//主函数-------------------------------------------

func main() {
	s := "abba"
	fmt.Println(longestPalindrome(s))
}
