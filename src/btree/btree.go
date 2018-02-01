package main

import (
	"fmt"
)

// 二叉树结构
type BTree struct {
	lchild *BTree
	value  int
	rchild *BTree
}


// 获取树的深度
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

/*
 *二叉树的遍历
 */

// 前序遍历
func (t *BTree) Preorder() {
	if t != nil {
		fmt.Print(t.value)
		t.lchild.Preorder()
		t.rchild.Preorder()
	}
}

// 中序遍历
func (t *BTree) Inorder() {
	if t != nil {
		t.lchild.Inorder()
		fmt.Print(t.value)
		t.rchild.Inorder()
	}
}

// 后序遍历
func (t *BTree) Postorder() {
	if t != nil {
		t.lchild.Postorder()
		t.rchild.Postorder()
		fmt.Print(t.value)
	}
}

// 层次遍历
//func (t *BTree) Levelorder() {
//	var q Queue
//	if t != nil {
//		q = q.Push(t)
//		for q.Len() != 0 {
//			res := q.Pop()
//			print(res.value)
//			if res.lchild != nil {
//				q = q.Push(res.lchild)
//			}
//			if res.rchild != nil {
//				q = q.Push(res.rchild)
//			}
//		}
//	}
//}

/*
 *二叉树构建
 */

// 构建搜索二叉树
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


// 序列化
func Serialize(A []int, head *BTree) {
	// pro-serialize

}

// 反序列化
func deserialize() {

}

// 通过前序后序、序列构建二叉树
// @preOrder前序序列
// @inOrder中序列
func Construct(preOrder []int, inOrder []int) *BTree {
	// ConstructCore函数的封装，只需用户提供序列就可以构建二叉树
	if preOrder == nil || inOrder == nil {
		return nil
	}
	return ConstructCore(preOrder, 0, len(preOrder)-1,
		inOrder, 0, len(inOrder)-1)
}

// 传入preOrder[]先序和中序inOrder []序列，用startPreOrder和endPreOrder标记起始位置和终止位置
func ConstructCore(preOrder []int, startPreOrder int, endPreOrder int,
	inOrder []int, startInOrder int, endInOrder int) *BTree {

	// 前序遍历序列的第一个数字是根结点的值
	rootValue := preOrder[startPreOrder]
	root := &BTree{nil, rootValue, nil}

	if startPreOrder == endPreOrder {
		if startInOrder == endInOrder &&
			preOrder[startPreOrder] == inOrder[startInOrder] {
			return root
		} else {
			fmt.Println("Invalid input!")
		}
	}

	// 在中序遍历中找到根结点的值
	rootInOrder := startInOrder
	for rootInOrder <= endInOrder && inOrder[rootInOrder] != rootValue {
		rootInOrder++
	}

	// 输入的两个序列不匹配的情况
	if rootInOrder == endInOrder && inOrder[rootInOrder] != rootValue {
		fmt.Println("Invalid input!")
	}

	leftLength := rootInOrder - startInOrder
	leftPreOrderEnd := startPreOrder + leftLength
	if leftLength > 0 {
		// 构建左子树
		root.lchild = ConstructCore(preOrder, startPreOrder+1, leftPreOrderEnd,
			inOrder, startInOrder, rootInOrder-1)
	}
	if leftLength < endPreOrder-startPreOrder {
		// 构建右子树
		root.rchild = ConstructCore(preOrder, leftPreOrderEnd+1, endPreOrder,
			inOrder, rootInOrder+1, endInOrder)
	}
	return root
}

func main() {
	//普通二叉树
	//              1
	//           /     \
	//          2       3
	//         /       / \
	//        4       5   6
	//         \         /
	//          7       8
	preorder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inorder := []int{4, 7, 2, 1, 5, 3, 8, 6}

	tree := Construct(preorder, inorder)
	tree.Inorder()
}
