package main

//【题目】
/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。

'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "*"
输出: true
解释: '*' 可以匹配任意字符串。
示例 3:

输入:
s = "cb"
p = "?a"
输出: false
解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
示例 4:

输入:
s = "adceb"
p = "*a*b"
输出: true
解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
示例 5:

输入:
s = "acdcb"
p = "a*c?b"
输入: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/wildcard-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
//
//【思路】
/*
?匹配规则简单，不用多说
重点是*的匹配，*匹配遵循最短匹配原则，即先只匹配一个字符
每次匹配*之后都要记录当前匹配时的状态（字符串指针和模式指针位置），以便于后续匹配失败时能回退进行更多的字符匹配
如果字符、？、*都无法匹配时则返回匹配失败

最后需要处理尾部情况
多个*存在时需要直接pass
然后看最终的指针长度能否和模式字符串长度匹配，可以则返回true
题目依然使用双指针来处理问题

*/

//【代码】
func isMatch(s string, p string) bool {
	var i, j, sIndex int
	var pIndex = -1
	for i < len(s) {

		if j < len(p) && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < len(p) && p[j] == '*' { //记录当前匹配位置
			pIndex = j
			sIndex = i
			j++
		} else if pIndex != -1 { //回退到上一个*匹配位置重新处理
			j = pIndex + 1
			sIndex++
			i = sIndex
		} else {
			return false
		}
	}
	for j < len(p) && p[j] == '*' { //处理尾部*字符
		j++
	}
	return j == len(p)
}

//【主函数】
func main() {

}

//【总结】：
/*







 */
