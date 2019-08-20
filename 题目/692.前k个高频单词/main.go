package main

import "fmt"

//【题目】
/*


 */
//
//【思路】
/*
1、创建一个结构体，包含单词以及单词个数
2、构建比较器，能够准确比较结构体的大小
3、在前k个数中维护一个k个数最小堆（找k个最大就维护最小堆，找k个最小就维护最大堆）
4、先heapify
5、遍历k之后的数，如果有比a[0]大的数，就交换至a[0],然后交换a[0]和a[k-1]，接着heapify（比a[0]小不合条件，直接跳过）
6、最后对k最小堆进行排序(a[0]和a[len(a)-1]交换，维护长度len(a)-1的堆)
*/

//【代码】
func topKFrequent(words []string, k int) []string {
	var m = map[string]int{}
	for _, word := range words {
		if _, ok := m[word]; !ok {
			m[word] = 1
		} else {
			m[word] += 1
		}
	}
	var list []word
	for key, value := range m {
		list = append(list, word{
			word:  key,
			count: value,
		})
	}
	Heapify(list[:k]) //在前k个数中维护一个最小堆

	for i := k; i < len(list); i++ {
		if list[i].LargeThan(list[0]) {
			list[i], list[0] = list[0], list[i]
			list[k-1], list[0] = list[0], list[k-1]
			Heapify(list[:k])
		}
	}
	fmt.Println(list)
	for i := k - 1; i > 0; i-- {
		list[0], list[i] = list[i], list[0]
		Heapify(list[:i])
	}

	var ret []string
	for i := 0; i < k; i++ {
		ret = append(ret, list[i].word)
	}
	return ret
}

//应该维护最小堆
func Heapify(a []word) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		var temp = i
		for j := temp*2 + 1; j < len(a); j = j*2 + 1 {
			if j+1 < len(a) && a[j].LargeThan(a[j+1]) {
				j++
			}
			if !a[temp].LargeThan(a[j]) {
				break
			}
			a[temp], a[j] = a[j], a[temp]
			temp = j
		}
	}
}
func (word1 word) LargeThan(word2 word) bool {
	if word1.count != word2.count {
		return word1.count > word2.count
	} else {
		return word1.word < word2.word //小单词更大
	}

}

type word struct {
	count int
	word  string
}

//【主函数】
func main() {
	var a = []string{"a", "aa", "aaa"}
	var k = 3
	b := topKFrequent(a, k)
	fmt.Println(b)
}

//【总结】：
/*







 */
