package main

//【题目】
/*
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-to-roman
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
// TODO
//【思路】
/*

可使用贪心算法将一些特殊的值列入map中,直接调用


*/

//【代码】
// 1
import (
	"fmt"
	"strings"
)

var RomanTable = [7]string{"I", "V", "X", "L", "C", "D", "M"}

func intToRoman(num int) string {
	// 题目规定 1<= num <=3999
	var romanTable []string
	var result = make([]string, 4)
	for i := 0; i < 4; i++ {
		romanTable = RomanTable[i*2:]
		switch num % 10 {
		case 1:
			result[3-i] = fmt.Sprint(romanTable[0])
		case 2:
			result[3-i] = fmt.Sprint(romanTable[0], romanTable[0])
		case 3:
			result[3-i] = fmt.Sprint(romanTable[0], romanTable[0], romanTable[0])
		case 4:
			result[3-i] = fmt.Sprint(romanTable[0], romanTable[1])
		case 5:
			result[3-i] = fmt.Sprint(romanTable[1])
		case 6:
			result[3-i] = fmt.Sprint(romanTable[1], romanTable[0])
		case 7:
			result[3-i] = fmt.Sprint(romanTable[1], romanTable[0], romanTable[0])
		case 8:
			result[3-i] = fmt.Sprint(romanTable[1], romanTable[0], romanTable[0], romanTable[0])
		case 9:
			result[3-i] = fmt.Sprint(romanTable[0], romanTable[2])
		}
		num /= 10
	}
	return strings.Join(result, "")
}

// 2
func IntToRoman(num int) string {
	ret := ""
	if num < 1 || num > 39999 {
		return ret
	}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	letters := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	// 重要思想用两个等长切片代替类python里的有序字典dict,go的map是无序的
	for i := 0; i < len(values); {
		if num >= values[i] { // num 逐次被分解
			num -= values[i]
			ret += letters[i]
		} else {
			i++
		}
	}
	return ret

}

//【主函数】
func main() {

}

//【总结】：
