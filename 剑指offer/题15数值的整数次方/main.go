package main

import "fmt"

// 实现Double Power(base , int)

func Power(base float64, n uint) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return base
	}
	result := Power(base, n>>1)
	result *= result
	if n&1 == 1 {
		result *= base
	}
	return result
}

func main() {
	fmt.Println(Power(2., 3))
}
