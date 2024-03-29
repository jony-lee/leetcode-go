package main

//【题目】
/*


 */
//
//【思路】
/*






 */

//【代码】
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var q = []*TreeNode{}
	var i, cur, next int //i层数、cur当前层元素个数、next下层元素个数
	var ret = [][]int{{}}
	q = append(q, root)
	cur = 1
	for len(q) != 0 {
		ret[i] = append(ret[i], q[0].Val)
		if q[0].Left != nil {
			q = append(q, q[0].Left)
			next++
		}
		if q[0].Right != nil {
			q = append(q, q[0].Right)
			next++
		}
		q = q[1:]
		cur--
		if cur == 0 && next != 0 {
			cur = next
			next = 0
			i++
			ret = append(ret, []int{})
		}
	}
	for i := 0; i < len(ret); i++ {
		if i%2 == 1 {
			reverse(ret[i])
		}
	}
	return ret
}
func reverse(a []int) {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

//【主函数】
func main() {

}

//【总结】：
/*







 */
