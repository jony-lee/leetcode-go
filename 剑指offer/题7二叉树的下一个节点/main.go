package main

// 给定一个二叉树和其中一个节点,如何找出中序遍历中的下一个节点?

// 思路
// 如果该节点有右子树,那么下一个节点就是右子树的最左子节点
// 如果该节点没有右子树:
// 		如果该节点是它父节点的左子节点,那么下一个节点就是它的父节点
// 		如果该一个节点既没有右子树,还是它父节点的右子节点
// 			则一直沿着父节点,向上遍历,直到找到一个是它父节点的左子节点的节点
//				如果该节点存在则它的父节点就是我们要找的节点

type treeNode struct {
	data           int
	lc, rc, parent *treeNode
}

func FindNextTreeNode(node *treeNode) *treeNode {
	p := node
	if p.rc != nil {
		p = p.rc
		for p.lc != nil {
			p = p.lc
		}
		return p
	} else if p.parent != nil {
		if p.parent.lc == p {
			return p.parent
		} else {
			for p.parent != nil && p.parent.lc != p {
				p = p.parent
			}
			return p.parent
		}
	}
	return nil
}
