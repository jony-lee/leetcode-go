package main

import "fmt"

// 输入一个整数n,输出该整数二进制中表示1的个数

// 方法1
// 先把n与1做 与运算,判断最低位是否是1(n & 1 = 0时最低位不为1),接着把1左移一位得到2,
// 在和n做与运算判断次最低位是否是1,循环

func NumOf1(n int) int {
	count := 0
	flag := 1
	for flag != 0 {
		if n&flag != 0 {
			count++
		}
		flag = flag << 1
	}
	return count
}

// 方法2
// 将一个整数和它减去1后的结果做与运算,会将它二进制中最右边的1变成0
func NumOF1(n int) int {
	count := 0
	for n != 0 {
		count++
		n = (n - 1) & n
	}
	return count
}

func main() {
	fmt.Println(NumOf1(8))
}
