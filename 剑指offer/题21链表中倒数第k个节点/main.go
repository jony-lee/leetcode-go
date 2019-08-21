package main

// 思路
// 第一个指针从链表头开始开始遍历,向前走k-1步,从k步开始第二个指针也开始从表头开始向前走
// 当第一个指针指向链表尾的时候,第二个指针正好指向倒数第k个节点

type ListNode struct {
	val  int
	next *ListNode
}

func findktotail(head *ListNode, k int) *ListNode {
	if k < 0 {
		return head
	}

	p1 := head
	p2 := head
	for i := 0; i < k-1; i++ {
		p1 = p1.next
	}
	for p1.next != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p2
}
