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

【解题思路】
遍历每一个元素，以当前元素为中心去寻找它的最大回文串，然后记录回文串长度，并找到最大值。
时间效率O(n2)
空间效率O()
总结：中心扩展法，理解起来比较简单，但是算法复杂度较高，马拉车算法时间复杂度低（O(n)），但理解有些困难，并且不具有泛用性
*/

//【代码】

func longestPalindrome(s string) string { // 中心扩展法
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

func max(a, b int) int {
	if a < b {
		max(b, a)
	}
	return a
}

func expandAroundCenter(s string, left int, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

//【主函数】

func main() {
	s := "abba"
	fmt.Println(longestPalindrome(s))
}
