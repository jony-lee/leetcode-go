package main

import "fmt"

// 打印1, 2, 3到最大的n位十进制数

func fun1(n int) {
	if n <= 0 {
		return
	}
	number := make([]rune, n)

	for i := 0; i < 0; i++ {
		number[0] = int32(i) + '0'
		printnum(number, n, 0)
	}
}

func printnum(n []rune, lens int, index int) {
	if index == lens-1 {
		Pnum(n)
		return
	}
	for i := 0; i < 10; i++ {
		n[index+1] = int32(i) + '0'
		printnum(n, lens, index+1)
	}
}

func Pnum(n []rune) {
	is0 := true
	lens := len(n)
	for i := 0; i < lens; i++ {
		if is0 && n[i] != '0' {
			is0 = false
		}
		if !is0 {
			fmt.Print(string(n[i]))
		}
	}
	fmt.Print('\t')
}

func main() {
	fun1(3)
}
