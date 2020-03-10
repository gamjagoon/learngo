package doublelink

import "fmt"

//Node : link
type Node struct {
	Next *Node
	Prev *Node
	Val  int
}

//DbLinkedList : goor
type DbLinkedList struct {
	Root *Node
	Tail *Node
}

//AddTail : find Tail and insert Val
func (l *DbLinkedList) AddTail(Val int) {
	if l.Root == nil {
		l.Root = &Node{Val: Val}
		l.Tail = l.Root
		return
	}

	l.Tail.Next = &Node{Val: Val}
	Prev := l.Tail
	l.Tail = l.Tail.Next
	l.Tail.Prev = Prev
}

//PrintNode : print all Val
func (l *DbLinkedList) PrintNode() {
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.Val)
}

//PrintRevers : print all Val
func (l *DbLinkedList) PrintRevers() {
	node := l.Tail
	for node.Prev != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.Prev
	}
	fmt.Printf("%d\n", node.Val)
}

//RemoveNode : Remove Node
func (l *DbLinkedList) RemoveNode(node *Node) {
	if node == nil {
		return
	}
	if node == l.Root {
		l.Root = l.Root.Next
		node.Next = nil
		l.Root.Prev = nil
		if l.Root == nil {
			l.Tail = nil
		}
		return
	}

	Prev := node.Prev

	if node == l.Tail {
		Prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = Prev
	} else {
		node.Prev = nil
		Prev.Next = Prev.Next.Next
		Prev.Next.Prev = Prev
	}
	node.Next = nil
}

//Find : find node
func (l *DbLinkedList) Find(Val int) *Node {
	node := l.Root
	for node != nil {
		if node.Val == Val {
			return node
		}
		node = node.Next
	}
	return node
}
