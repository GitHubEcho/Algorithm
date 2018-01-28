package main

import (
	btree "btree"
)

func main() {
	// Dserializaton for create tree
	t := btree.New(6, 1)
	//fmt.Println(t1.getDepth())
	//t1.preorder()

	//t1.inorder()

	t.Levelorder()  //层次遍历

	//t1.postorder()
}

