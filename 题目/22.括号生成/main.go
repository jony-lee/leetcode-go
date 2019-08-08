package main

//【题目】
/*
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
// TODO
//【思路】
/*

向一个空字符串cur逐一添加"("和")"
添加规则是:
"("数量小于n时可添加"("
")"数量小于"("数量时可以添加")"
重复上述规则
直到len(cur) == n *2 为止

*/

//【代码】
func generateParenthesis(n int) []string {
	ans := []string{} // 记录所有可能
	backtrack(&ans, "", 0, 0, n)
	return ans
}

// 不断的向cur添加合理的字符,直到len(cur) == max *2 为止
func backtrack(ans *[]string, cur string, open int, close int, max int) {
	if len(cur) == max*2 { // 输出所有情况
		*ans = append(*ans, cur)
		return
	}
	if open < max { // "("数量小于n时可添加"("
		backtrack(ans, cur+"(", open+1, close, max)
	}
	if close < open { // ")"数量小于"("数量时可以添加")"
		backtrack(ans, cur+")", open, close+1, max)
	}
}

//【主函数】
func main() {

}

//【总结】：
/*







 */
