/*
 * Author: skyler
 * Date: 2018-01-30
 */

/*重建二叉树
 *描述：
 * 输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。
 *  假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
 *  例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
 *分析：
 *  先根据前序遍历序列的第一个数字创建根结点，接下来在中序遍历序列中找到根结点的位置，这样就能确定左、右子树结点的数量。
 * 在前序遍历和中序遍历的序列中划分了左、右子树结点的值之后，就可以递归地去分别构建它的左右子树。
 */

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

// 后序遍历
func (t *BTree) Inorder() {
	if t != nil {
		t.lchild.Inorder()
		fmt.Print(t.value)
		t.rchild.Inorder()
	}
}

// 通过前序后序、序列构建二叉树
// @preOrder前序序列
// @inOrder中序列
func Construct(preOrder []int, inOrder []int ) *BTree {
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

// Serialization
func Serialize(A []int, head *BTree) {
	// pro-serialize

}

//Deserialization
func deserialize() {

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
