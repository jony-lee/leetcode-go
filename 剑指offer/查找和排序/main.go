package main

import (
	"fmt"
)

// 快速排序
// 先在数组中选择一个数字n,接下来把数组中比n小的数字移到数组左边,比n大数字移到数组右边

func main() {
	a := []int{3, 5, 8, 1, 2, 9, 4, 7, 6}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)

}
