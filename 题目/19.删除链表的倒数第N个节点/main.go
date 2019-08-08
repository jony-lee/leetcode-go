package main

//【题目】
/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
// TODO
//【思路】
/*

双指针发和两次迭代法


*/

//【代码】

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := &ListNode{Next: head} // 建立指向第一个元素的头结点
	fast, slow, step := node, node, 0
	for step < n { // fast向前移动n次
		fast = fast.Next
		step++
	}
	for fast.Next != nil { // fast和slow同时向前移动,直到fast指向最后一个
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return node.Next
}

// 两次迭代法
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	node := &ListNode{Next: head} // 用go实现链表一定要创建头结点,以免发生无效指针的情况
	pointer, pointer2, length := node, node, 1
	for pointer.Next != nil {
		pointer = pointer.Next
		length++
	}
	index := length - n
	for i := 1; i <= length; i++ { // 向前移动 length
		if i == index {
			if pointer2.Next == nil {
				break
			}
			pointer2.Next = pointer2.Next.Next
			break
		} else {
			pointer2 = pointer2.Next
		}
	}
	return node.Next
}

// 错误案例 // 没有考虑到n = len的情况即删除第一个点
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	h := head
	l := 1
	for h != nil {
		h = h.Next
		l++
	}
	l = l - n - 1
	h = head
	for i := 0; i < l-1; i++ { // 向前移动l-1
		h = h.Next
	}
	if n == 1 {
		h.Next = nil
		return head
	}
	h.Next = h.Next.Next
	return head
}

//【主函数】
func main() {

}

//【总结】：
/*







 */
