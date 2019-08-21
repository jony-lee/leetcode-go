package main

// 在二维数组中,每一行每一列都是顺序排序,输入一个整数和这样的二维数组,判断数组中是否存在该整数

// 思路
// 先取数组中右上角的数字,如果该数字等于要查找的数字,则查找过程结束
// 如果该数大于要查找的数字,则剔除这个数字所在的行
// 如果该数字小于要查找的数字,则剔除这个数字所在的列

func Find(nums [][]int, target int) bool {
	for i, j := 0, len(nums)-1; i < len(nums) && j >= 0; {
		if nums[i][j] == target {
			return true
		} else if nums[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
