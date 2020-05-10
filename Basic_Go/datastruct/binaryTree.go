package datastruct

import "fmt"

//BinaryTreeNode : node of BST
type BinaryTreeNode struct {
	Val         int
	left, right *BinaryTreeNode
}

//BinaryTree :
type BinaryTree struct {
	Root *BinaryTreeNode
}

type depthNode struct {
	depth int
	node  *BinaryTreeNode
}

//NewBinaryTree : make New Tree
func NewBinaryTree(v int) *BinaryTree {
	tree := &BinaryTree{}
	tree.Root = &BinaryTreeNode{Val: v}
	return tree
}

//AddNode : add value
func (n *BinaryTreeNode) AddNode(v int) *BinaryTreeNode {
	if n.Val > v {
		if n.left == nil {
			n.left = &BinaryTreeNode{Val: v}
			return n.left
		}
		return n.left.AddNode(v)
	}
	if n.right == nil {
		n.right = &BinaryTreeNode{Val: v}
		return n.right
	}
	return n.right.AddNode(v)
}

//Print : print BST
func (t *BinaryTree) Print() {
	q := []*depthNode{}
	q = append(q, &depthNode{depth: 0, node: t.Root})
	currentDepth := 0

	for len(q) > 0 {
		var first *depthNode
		first, q = q[0], q[1:]

		if first.depth != currentDepth {
			fmt.Println()
			currentDepth = first.depth
		}
		fmt.Printf("%d ", first.node.Val)

		if first.node.left != nil {
			q = append(q, &depthNode{depth: currentDepth + 1, node: first.node.left})
		}

		if first.node.right != nil {
			q = append(q, &depthNode{depth: currentDepth + 1, node: first.node.right})
		}
	}
}

//Search : tree for BinaryTree
func (t *BinaryTree) Search(v int) (bool, int) {
	return t.Root.Search(v, 0)
}

//Search : find value
func (n *BinaryTreeNode) Search(v int, cnt int) (bool, int) {

	if n.Val == v {
		return true, cnt
	}
	if n.Val > v {
		if n.left != nil {
			return n.left.Search(v, cnt+1)
		}
		return false, cnt
	}
	if n.right != nil {
		return n.right.Search(v, cnt+1)
	}
	return false, cnt
}
