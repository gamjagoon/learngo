package datastruct

import "fmt"

//TreeNode : tree
type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

//Tree : struct tree
type Tree struct {
	Root *TreeNode
}

//AddNode : add node to TreeNode
func (t *TreeNode) AddNode(val int) {
	t.Childs = append(t.Childs, &TreeNode{Val: val})
}

//AddNode : add node
func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}

//Dfs1 : print node
func (t *Tree) Dfs1(node *TreeNode) {
	fmt.Printf("%d->", node.Val)
	for i := 0; i < len(node.Childs); i++ {
		t.Dfs1(node.Childs[i])
	}
}

//Dfs2 :dasdf
func (t *Tree) Dfs2() {
	s := []*TreeNode{}
	s = append(s, t.Root)

	for len(s) > 0 {
		var top *TreeNode
		top, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d->", top.Val)
		for i := 0; i < len(top.Childs); i++ {
			s = append(s, top.Childs[i])
		}
	}
}
