package main

//【题目】
/*


 */
//
//【思路】
/*
层次遍历，维护两个变量，一个变量存储队列中当前层次的元素个数，另一个变量存储下一层存储的元素个数
如果当前层元素个数为零了，就让把下一层的元素个数赋给当前层，下一层元素个数清零
这样就能够判断出某一层是否访问完成。





*/

//【代码】
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var q = []*TreeNode{}
	var cur, next int
	var target = &TreeNode{}
	q = append(q, root)
	cur = 1
	for len(q) != 0 {
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
		if cur == 0 && next != 0 { //当前层遍历完
			cur = next
			next = 0
			if len(q) != 0 {
				target = q[0]
			}
		}
	}
	return target.Val
}

func findBottomLeftValue1(root *TreeNode) int {
	data := make([][]int, 0)
	data = fun1(data, 0, root)
	return data[len(data)-1][0]
}

func fun1(data [][]int, i int, node *TreeNode) [][]int {
	if len(data)-1 < i {
		data = append(data, []int{})
	}
	data[i] = append(data[i], node.Val)
	if node.Left != nil {
		data = fun1(data, i+1, node.Left)
	}
	if node.Right != nil {
		data = fun1(data, i+1, node.Right)
	}
	return data
}

//【主函数】
func main() {

}

//【总结】：
/*







 */
