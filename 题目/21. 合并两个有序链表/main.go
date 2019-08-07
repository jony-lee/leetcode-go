package main

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
