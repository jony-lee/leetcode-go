package main

// 输入两颗二叉树A, B.判断B是不是A的子结构

// 思路
// 第一步在A中找到和B根节点一样值的节点R
// 第二步,判断树A中,以R为根节点的子树是不是包含和B一样的结构

type Node struct {
	value int
	left  *Node
	right *Node
}

func HasSubTree(A, B *Node) bool { // 第一步
	result := false
	if A != nil && B != nil {
		if A.value == B.value {
			result = Dose(A, B)
		}
		if !result {
			result = HasSubTree(A.left, B)
		}
		if !result {
			result = HasSubTree(A.right, B)
		}
	}
	return result
}

func Dose(A, B *Node) bool { // 第二步
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.value != B.value {
		return false
	}
	return Dose(A.left, B.left) && Dose(A.right, B.right)
}
