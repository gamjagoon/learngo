package main

import (
	"fmt"

	"github.com/gamjagoon/Tucker/datastruct"
)

func main() {
	tree := datastruct.Tree{}

	val := 1
	tree.AddNode(val)
	val++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}
	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(val)
			val++
		}
	}

	tree.Dfs1(tree.Root)
	fmt.Println()
	tree.Dfs2()
}
