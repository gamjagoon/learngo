package main

import (
	"fmt"
	"math/rand"

	"github.com/gamjagoon/learngo/Tucker/datastruct"
)

func main() {

	tree := datastruct.NewBinaryTree(50)
	for i := 0; i < 100; i++ {
		tree.Root.AddNode(rand.Intn(100))
	}
	tree.Print()
	fmt.Println()
	findnum := rand.Intn(100)
	if found, cnt := tree.Search(findnum); found {
		fmt.Println("found ", findnum, "cnt : ", cnt)
	} else {
		fmt.Println("not found ", findnum, "cnt : ", cnt)
	}
}
