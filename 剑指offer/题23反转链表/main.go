package main

type ListNode struct {
	val  int
	next *ListNode
}

func reverseList(node *ListNode) *ListNode {
	var PRH *ListNode
	var pNode = node
	var pPre *ListNode
	for pNode != nil {
		pNext := pNode.next
		if pNext != nil { // 如果下一节点非空
			PRH = pNode
		}
		pNode.next = pPre
		pPre = pNode
		pNode = pNext
	}
	return PRH
}
