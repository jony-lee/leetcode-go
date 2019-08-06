package main

import "fmt"

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


解题思路：
找中位数，第一个想法是遍历一遍，但是这个算法复杂度就是O(m+n)了，似乎不符合题意啊
先求出两数组m、n的中位数a、b，如果b>a则结果应该在a后面数组或b前面数组中，那么a前面数组和b后面数组元素可以舍弃。
//TODO

时间效率O()
空间效率O()
总结：
*/

//思路1-------------------------------------------

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//len1, len2 := len(nums1), len(nums2)
	//var low, high []int
	//var j,k int
	//if nums1[0]>nums2[0]{
	//	low= nums2[:]
	//}else {
	//	low=nums1[:]
	//}
	//if nums1[len1-1]>nums2[len2-1]{
	//	high= nums1[:]
	//}else {
	//	high=nums2[:]
	//}
	//start,end:= 0,len(high)-1
	//for{
	//	if start>len(low)-1{
	//		break
	//	}
	//	if end<0{
	//		break
	//	}
	//	j = (len(low)-1)/2	//中位数或中位数左边一个
	//	k = (len(high))/2	//中位数或中位数右边一个
	//	if low[j] > high[0]{
	//		low = low[j:]
	//	}else{
	//
	//	}
	//}
	//if

	//len1, len2 := len(nums1), len(nums2)
	//k := (len1 + len2) / 2 //k为奇数，则正好是中位数，k为偶数，则需要和临近左边一个数相加除2
	//if nums1[0] > nums2[len2-1] {
	//	if k < len1 {
	//		return nums1[k]
	//	}
	//}
	//
	//a := len1 / 2
	//b := len2 / 2

}

//主函数-------------------------------------------

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a[len(a)-1])
}
