package main

import "fmt"

//【题目】
/*
报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。



示例 1:

输入: 1
输出: "1"
示例 2:

输入: 4
输出: "1211"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-and-say
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
//
//【思路】
/*
定义genNext函数，根据上一个字符串内容返回字符串内容
迭代计算即可
注意边界条件：第一个生成的是"1"




*/

//【代码】
func countAndSay(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s = genNext(s)
	}
	return s
}
func genNext(s string) string {
	if len(s) == 0 {
		return "1"
	}
	if len(s) == 1 {
		return "1" + s
	}
	s2 := ""
	temp := s[0]
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] != temp {
			s2 += string(count+'0') + string(temp)
			temp = s[i]
			count = 1
		} else {
			count++
		}
	}
	//加最后一个数
	s2 += string(count+'0') + string(temp)
	return s2
}

//【主函数】
func main() {
	fmt.Println(countAndSay(1))
}

//【总结】：
/*







 */
