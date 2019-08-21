package main

// 思路使用递归会比较简单

type ListNode struct {
	val  int
	next *ListNode
}

func merge(h1 *ListNode, h2 *ListNode) *ListNode {
	if h1 == nil { // 重点
		return h2
	} else if h2 == nil {
		return h1
	}
	var pn *ListNode

	if h1.val < h2.val {
		pn = h1
		pn.next = merge(h1.next, h2)
	} else {
		pn = h2
		pn.next = merge(h1, h2.next)
	}
	return pn
}
