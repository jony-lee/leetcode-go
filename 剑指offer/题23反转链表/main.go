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
		pNext := pNode.next // 记录下一节点
		if pNext != nil {   // 如果下一节点非空
			PRH = pNode // 新反转链表头指向这个节点
		}
		pNode.next = pPre // 将本节点指向他的前一个节点,(第一个节点指向nil)
		pPre = pNode      // 前一个节点更新为当前结点
		pNode = pNext
	}
	return PRH
}
