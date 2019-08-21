package main

import "fmt"

// 实现一个函数用来判断字符串是否表示一个数值

func isNum(str string) bool {
	if str == "" {
		return false
	}
	numeric, i := sancInt(str, 0)
	if i < len(str) && str[i] == '.' {
		i++
		a := false
		a, i = sancUint(str, i) // 判断小数部分
		numeric = a || numeric
	}
	if i < len(str) && (str[i] == 'e' || str[i] == 'E') {
		i++
		a, _ := sancInt(str, i)
		numeric = numeric && a
	}
	return numeric
}

func sancUint(str string, i int) (bool, int) {
	n := i
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		i++
	}

	return i > n, i
}

func sancInt(str string, i int) (bool, int) {
	if str[i] == '+' || str[i] == '-' {
		i++
	}
	return sancUint(str, i)
}

func main() {
	fmt.Print(isNum("123e-121"))
}
