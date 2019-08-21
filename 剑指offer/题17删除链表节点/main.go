package main

type ListNode struct {
	val  int
	next *ListNode
}

// 把要删除的节点的下一个节点信息复制到要删除的节点上

func DeleteNode(head **ListNode, d *ListNode) {
	if d.next != nil {
		p := d.next
		d.val = p.val
		d.next = p.next
		p = nil
	} else if *head == d {
		*head = nil
	} else {
		p := *head
		for p.next != d {
			p = p.next
		}
		p.next = nil
	}
}
