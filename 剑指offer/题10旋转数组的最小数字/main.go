package main

// 吧一个数组的开始的若干个元素搬到数组的末尾,称之为旋转
// 输入一个递增排序的数组的一个旋转,输出旋转数组的最小元素

func Min(d []int) int {
	a, b := 0, len(d)-1
	p := (a + b) >> 1
	for a < b {
		if d[p] > d[b] {
			a = p + 1
		} else {
			b = p
		}
		p = (a + b) >> 1
	}
	return d[p]
}

func main() {
	a := []int{0, 1, 2}
	println(Min(a))
}
