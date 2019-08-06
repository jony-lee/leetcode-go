package main

import (
	"fmt"
)

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


解题思路：
最开始想到暴力匹配，遍历所有子串，再求出最大子串长度，但是这样结果过于庞大，实现上就不考虑了。
然后想到用堆结构，通过堆结构来找到最长的子串，不过好像用不上。
看解题思路中说主要是靠滑动窗口算法。
大体思路理解了，就是创建一个队列，一个个元素进队列，如果发生元素重复了，就移除左边队列重复的值，这样就能保证队列内元素永远是不重复的
也就是队列即非重复子串，然后每移动一次都记录下队列长度，最后给出最大的队列长度。
标记子字符串首、尾 start和end，标记当前字符串长度和最大字符串长度length和max
遍历字符v，检查v在map中是否重复，如果有重复就移动start到老位置向右一个
都需要移动end标签和更新v的位置
最后将当前字符串长度与最大字符串长度进行比较即可
一次遍历即可找出无重复最长子串
时间效率O(n)
空间效率O()
总结：这道题最开始没想到用map来解决，后来看了别人代码，唉，发现查找问题，还是首选考虑用上map
另外就是这个题目可以做一些拓展，比如获取最长子串，只需要在下面代码加入一个字符串变量，实时记录最长的那个字符串就可以
*/

//思路1-------------------------------------------

func lengthOfLongestSubstring(s string) int {
	start, end, length, max := 0, 0, 0, 0
	m := make(map[int32]int)
	for i, v := range s {

		if old, ok := m[v]; ok == true && start < old+1 {
			start = old + 1
		}
		end = i
		m[v] = i
		length = end - start + 1
		if max < length {
			max = length
		}
	}
	return max
}

//主函数-------------------------------------------

func main() {
	s := "abba"
	fmt.Println(lengthOfLongestSubstring(s))
}
