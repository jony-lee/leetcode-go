package main

import "fmt"

// 写一个函数,输入n,求斐波那契数列的第n项

func Fibonacci(n int) int {
	result := []int{0, 1}
	if n < 2 {
		return result[n]
	}
	f1 := 1
	f2 := 2
	fn := 0
	for i := 2; i < n; i++ {
		fn = f1 + f2
		f1 = f2
		f2 = fn
	}
	return fn
}

func fib(n int) int {
	results := make([]int, n+1)
	results[0] = 0
	results[1] = 1
	results[2] = 2
	for i := 3; i <= n; i++ {
		results[i] = results[i-1] + results[i-2]
	}
	return results[n]
}

// 青蛙跳台阶问题
// 一只青蛙一次可以跳上1级台阶,也可以跳上2级台阶,求青蛙跳上一个n级台阶有多少种跳法?
// 跳法数目f(n)等于后面剩下的n-1级台阶的跳法数目f(n-1)加上后面剩下n-2级台阶跳法数目f(n-2)之和
// 故相同的算法

func main() {
	fmt.Println(fib(8))
	fmt.Println(Fibonacci(8))
}
