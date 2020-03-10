package datastruct

//HeapTreeNode : Node in HeapTree
type HeapTreeNode struct {
	val         int
	left, right *HeapTreeNode
}

//HeapTree : true is maxheap, false is minheap
type HeapTree struct {
	maxheap bool
	Root    *HeapTreeNode
}

//NewHeapTree : make new heaptree
func NewHeapTree(max bool, value int) *HeapTree {
	return &HeapTree{maxheap: max, Root: &HeapTreeNode{val: value}}
}

//Push : add node
func (t *HeapTree) Push(val int) {
	if t.maxheap {
		t.Root.Pushmax()
	} else {
		t.Root.Pushmine()
	}
}

//Pushmax : make maxheap
func (t *HeapTreeNode) Pushmax(val int) *HeapTreeNode {
	if t.left.val > val {

	}
}
