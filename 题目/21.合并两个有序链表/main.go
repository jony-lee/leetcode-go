package main

//【题目】
/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
// TODO
//【思路】
/*






 */

//【代码】
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	node1 := l1
	node2 := l2
	node3 := &ListNode{}
	p := node3
	for node2 != nil && node1 != nil {
		if node1.Val < node2.Val {
			node3.Next = node1
			node1 = node1.Next
		} else {
			node3.Next = node2
			node2 = node2.Next
		}
		node3 = node3.Next
	}
	if node2 == nil {
		node3.Next = node1
	} else {
		node3.Next = node2
	}
	return p.Next

}

//【主函数】
func main() {

}

//【总结】：
/*







 */
