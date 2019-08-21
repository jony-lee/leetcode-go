package main

// 首先用快慢指针确定换中的某一节点位置,
// 从这个节点出发,继续向前,并不断计数,若又返回到了这个节点
// 则得到了换内节点个数n
// 设置指针p1,p2,开始均指向表头
// 先让p1前进n,再p1, p2同时向前,当p1 == p2时,则p1指向了入环的入口

type ListNode struct {
	val  int
	next *ListNode
}

func meetnode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ps := head.next
	if ps == nil {
		return nil
	}
	pf := head.next
	for pf != nil && ps != nil {
		if pf == ps {
			return pf
		}
		ps = ps.next
		pf = pf.next
		if pf != nil {
			pf = pf.next
		}
	}
	return nil
}

func entryNodeOfloop(head *ListNode) *ListNode {
	meetnode := meetnode(head)
	if meetnode == nil {
		return nil
	}
	nodesINloop := 0
	p1 := meetnode
	for p1.next != meetnode {
		p1 = p1.next
		nodesINloop++
	}
	p1 = head
	for i := 0; i < nodesINloop; i++ {
		p1 = p1.next
	}
	p2 := head
	for p1 != p2 {
		p1 = p1.next
		p2 = p2.next
	}
	return p1
}
