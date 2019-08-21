package main

// 使用两个栈实现一个队列,并实现队列的入队和出队操作

// 设stack1 stack2
// 入队:将元素直接stack1.push
// 出队:如果stack2不为空,则直接stack2.pop
//			否则将stack1中元素全部pop出,并且压入stack2中,再从stack2弹出元素

import "fmt"

type stack struct {
	val  interface{}
	next *stack
}

func (l *stack) Init(a []int) {
	l.val = len(a)
	p := l
	for i := 0; i < len(a); i++ {
		p.next = &stack{
			val:  a[i],
			next: nil,
		}
		p = p.next
	}
}

func (l stack) String() string {
	p := &l
	s := fmt.Sprintf("stack: (len-%d) @ ", l.val)
	for p.next != nil {
		s += fmt.Sprintf("-> %d ", p.next.val)
		p = p.next
	}
	return s
}

func (l *stack) Pop() interface{} {
	p := l
	val := p.next.val
	p.next = p.next.next
	l.val = l.val.(int) - 1
	return val
}

func (l *stack) Push(e interface{}) {
	p := l
	p = p.next
	a := stack{
		val:  e,
		next: p,
	}
	l.next = &a
	l.val = l.val.(int) + 1
}

func (l *stack) Empty() bool {
	if l.val.(int) > 0 {
		return false
	} else {
		return true
	}
}

type queue struct{}

var stack1 stack
var stack2 stack

func (q *queue) push(data interface{}) {
	stack1.Push(data)
}

func (q *queue) pop() interface{} {
	if !stack2.Empty() {
		return stack2.Pop()
	} else {
		for !stack1.Empty() {
			stack2.Push(stack1.Pop())
		}
	}
	if stack2.Empty() {
		return nil
	} else {
		return stack2.Pop()
	}
}

func main() {
	stack1.Init([]int{})
	stack2.Init([]int{})
}
