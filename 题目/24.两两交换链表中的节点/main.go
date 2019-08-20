package main

import (
	"fmt"
)

//【题目】
/*


 */
// TODO
//【思路】
/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

//【代码】

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	h := &ListNode{Next: head}
	r := h
	p1 := head
	if h.Next == nil {
		return h.Next
	}
	p2 := p1.Next
	for p2 != nil { // 奇数情况是否到最后
		r.Next = p2
		p1.Next = p2.Next
		p2.Next = p1
		r = p1
		p1 = p1.Next
		if p1 == nil { // 偶数情况是否到最后
			break
		}
		p2 = p1.Next
	}
	return h.Next
}

//【主函数】
func main() {
	test := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	a := reverseKGroup(test, 2)
	fmt.Println(a)
}

func reverseKGroup1(head *ListNode, k int) *ListNode { // 翻转题优先考虑栈
	if head == nil {
		return head
	}
	N_new := &ListNode{}
	h := N_new
	p := &ListNode{Next: head} // 头结点
	p1 := p
	a := []*ListNode{}
	for j := 1; p.Next != nil; j++ {
		p = p.Next              // 向前进位,当前为第j个元素
		a = append(a, p)        // 加入堆中
		if j%k == 0 && j != 0 { // 当j为k的整数倍时,出堆
			p1 = p
			for i := len(a) - 1; i >= 0; i-- {
				N_new.Next = &ListNode{Val: a[i].Val}
				N_new = N_new.Next // 指向最后一个操作结点
			}
			a = []*ListNode{}
		}
	}
	N_new.Next = p1.Next
	return h.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	tail := head
	listLen := 1
	for i := 0; i < k-1; i++ {
		if tail.Next == nil {
			break
		}
		tail = tail.Next
		listLen++
	}
	if listLen < k {
		return head
	}
	next := tail.Next
	head, tail = reverseList(head, tail)
	if next != nil {
		tail.Next = reverseKGroup(next, k)
	}
	return head
}

//【总结】：
/*







 */
