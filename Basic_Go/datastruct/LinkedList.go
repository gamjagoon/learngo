package datastruct

import "fmt"

//Node : link
type Node struct {
	Next *Node
	Val  int
}

//Linklist : goor
type Linklist struct {
	Root *Node
	Tail *Node
}

//AddTail : find Tail and insert Val
func (l *Linklist) AddTail(Val int) {
	if l.Root == nil {
		l.Root = &Node{Val: Val}
		l.Tail = l.Root
		return
	}

	l.Tail.Next = &Node{Val: Val}
	l.Tail = l.Tail.Next
}

//PrintNode : print all Val
func (l *Linklist) PrintNode() {
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.Val)
}

//RemoveNode : Remove Node
func (l *Linklist) RemoveNode(node *Node) {
	if node == nil {
		return
	}
	if node == l.Root {
		l.Root = l.Root.Next
		node.Next = nil
		if l.Root == nil {
			l.Tail = nil
		}
		return
	}

	prev := l.Root
	for prev.Next != node {
		prev = prev.Next
	}

	if node == l.Tail {
		prev.Next = nil
		l.Tail = prev
	} else {
		prev.Next = prev.Next.Next
	}
	node.Next = nil
}

//Find : find node
func (l *Linklist) Find(Val int) *Node {
	node := l.Root
	for node != nil {
		if node.Val == Val {
			return node
		}
		node = node.Next
	}
	return node
}

//Back : retrun val
func (l *Linklist) Back() int {
	return l.Tail.Val
}

//Removelast :
func (l *Linklist) Removelast() {
	if l.Tail == nil {
		return
	}
	l.RemoveNode(l.Tail)
}
