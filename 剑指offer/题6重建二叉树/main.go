package main

// 输入某二叉树前序遍历和中序遍历的结果,请重建二叉树

// 思路
// 前序遍历的第一个是root节点
// 确定root节点后,在中序遍历中确定左右子树
// 在前序遍历中确定左右子树的根
// 循环上述步骤

import (
	"fmt"
)

// 重建二叉树

type Node struct {
	value int
	left  *Node
	right *Node
}

func main() {
	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inOrder := []int{4, 7, 2, 1, 5, 3, 8, 6}
	tree := constructBTree(preOrder, inOrder)

	//将构建好的二叉树 输出先序遍历和中序遍历的结果 用于检验
	preCatTree(tree)
	fmt.Println()
	inCatTree(tree)
}

//重建二叉树
func constructBTree(preOrder, inOrder []int) *Node {
	l := len(preOrder)
	if l == 0 {
		return nil
	} // 输出条件
	root := &Node{ // 建立根节点
		value: preOrder[0],
	}
	if l == 1 {
		return root
	}
	index := 0
	for i, v := range inOrder {
		if v == root.value {
			index = i
		}
	}
	fmt.Println("左子树", preOrder[1:index+1], inOrder[0:index]) //可打开注释查看详细过程
	root.left = constructBTree(preOrder[1:index+1], inOrder[0:index])
	fmt.Println("右子树", preOrder[index+1:], inOrder[index+1:])
	root.right = constructBTree(preOrder[index+1:], inOrder[index+1:])
	return root
}

func preCatTree(t *Node) {
	fmt.Print(t.value, " ")
	if t.left != nil {
		preCatTree(t.left)
	}
	if t.right != nil {
		preCatTree(t.right)
	}
}
func inCatTree(t *Node) {
	if t.left != nil {
		inCatTree(t.left)
	}
	fmt.Print(t.value, " ")
	if t.right != nil {
		inCatTree(t.right)
	}
}
