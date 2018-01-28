package btree


import (
	"math/rand"
	"fmt"
)

// binary tree implement
type BTree struct {
	lchild *BTree
	value  int
	rchild *BTree
}

func New(n, k int) *BTree {
	var t *BTree
	for _, v := range rand.Perm(n) {
		//fmt.Print(v)
		t = t.insert(v * k)
	}
	fmt.Print("\n")
	return t
}

func (t *BTree) insert(v int) *BTree {
	if t == nil {
		return &BTree{nil, v, nil}
	}
	if v < t.value {
		t.lchild = t.lchild.insert(v)
		return t
	}
	t.rchild = t.rchild.insert(v)
	return t
}

// get tree depth
func (t *BTree) GetDepth() int {
	var rd int
	var ld int
	if t == nil {
		return 0;
	} else {
		ld = t.lchild.GetDepth()
		rd = t.rchild.GetDepth()
		if ld > rd {
			return ld + 1
		}
		return rd + 1
	}
}

//travail tree
func (t *BTree) Preorder() {
	if t != nil {
		fmt.Print(t.value)
		t.lchild.Preorder()
		t.rchild.Preorder()
	}
}

func (t *BTree) Inorder() {
	if t != nil {
		t.lchild.Inorder()
		fmt.Print(t.value)
		t.rchild.Inorder()
	}
}

func (t *BTree) Postorder() {
	if t != nil {
		t.lchild.Postorder()
		t.rchild.Postorder()
		fmt.Print(t.value)
	}
}

// level traversal
func (t *BTree)Levelorder()  {
	var q Queue
	if t != nil{
		q = q.Push(t)
		for q.Len() != 0 {
			res := q.Pop()
			print(res.value)
			if res.lchild != nil {
				q = q.Push(res.lchild)
			}
			if res.rchild != nil {
				q = q.Push( res.rchild)
			}
		}
	}
}

// Serialization
func Serialize(A []int ,head *BTree){
	// pro-serialize

}

//Deserialization
func deserialize()  {

}

