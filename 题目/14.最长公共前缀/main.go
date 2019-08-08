package main

//【题目】
/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
// TODO
//【思路】
/*


按列搜索



*/

//【代码】
func longestCommonPrefix(s []string) string {
	temp := ""
	if len(s) == 0 { // 特殊情况
		return ""
	}
	minlen := len(s[0])
	for _, j := range s { // 求最短字符串的长度
		if minlen > len(j) {
			minlen = len(j)
		}
	}
	for i := 0; i < minlen; i++ { // 按列i搜索
		a := s[0][i]                  // 初始化为第一个字符串中的第i个字符
		for j := 0; j < len(s); j++ { // 遍历所有字符串中第i个字符
			if a != s[j][i] { // 如果出现不同则返回
				return temp
			}
		}
		temp += string(a) // 否则加入最长公共前缀的尾部
	}
	return temp
}

//【主函数】
func main() {

}

//【总结】：
/*







 */
