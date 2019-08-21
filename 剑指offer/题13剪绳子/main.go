package main

import "fmt"

// 给一个长度为n的绳子,把绳子剪成m段 m>1 n>1
// 每段绳子的长度的最大乘积是多少?

// 动态规划
// 定义函数f(n)为把长度为n的绳子剪成若干段后各段长度乘积的最大值,在剪下第一刀的时候有n-1种选择
// 所以f(n) = max(f(i)*f(n-i))

func maxP(lens int) int {
	if lens < 2 {
		return 0
	} else if lens == 2 {
		return 1
	} else if lens == 3 {
		return 2
	}
	products := make([]int, lens+1)
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3 // 从第4开始,最大乘积才能超过他本身的长度
	max := 0
	for i := 4; i <= lens; i++ {
		max = 0
		for j := 1; j <= i/2; j++ {
			product := products[j] * products[i-j]
			if max < product {
				max = product
			}
		}
		products[i] = max
	}
	max = products[lens]
	return max
}

// 贪婪算法
// 当n>=5时,我们尽可能多减长度为3的绳子,剩下的长度为4时,把绳子剪成两端长度为2的绳子
func maxP2(lens int) int {
	if lens < 2 {
		return 0
	} else if lens == 2 {
		return 1
	} else if lens == 3 {
		return 2
	}
	timeOf3 := lens / 3
	if lens-timeOf3*3 == 1 { // 当剩下长度为4时
		timeOf3 -= 1
	}
	timeOf2 := (lens - timeOf3*3) / 2
	return pow(3, timeOf3) * pow(2, timeOf2)
}

func pow(n, m int) int {
	se := 1
	for i := 0; i < m; i++ {
		se *= n
	}
	return se
}
func main() {
	fmt.Println(maxP(8))
	fmt.Println(maxP2(8))
}
