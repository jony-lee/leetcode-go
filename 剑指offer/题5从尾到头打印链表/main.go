package main

import "fmt"

type stack struct {
	val  int
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
	s := fmt.Sprintf("List: (len-%d) @ ", l.val)
	for p.next != nil {
		s += fmt.Sprintf("-> %d ", p.next.val)
		p = p.next
	}
	return s
}

func (l *stack) Empty() bool {
	if l.val > 0 {
		return false
	} else {
		return true
	}
}

// 思路反序优先想到栈,想到栈就想到递归

func (l *stack) PrintListReversingly() {
	if l != nil {
		if l.next != nil {
			l.next.PrintListReversingly()
		}
		fmt.Print(l.val, " <- ")
	}
}

func main() {
	var list stack
	list.Init([]int{1, 2, 3, 3, 5, 6})
	list.PrintListReversingly()
}
