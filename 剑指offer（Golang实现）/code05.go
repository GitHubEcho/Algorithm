/*
 Author: skyler
 Date: 2018-01-29
*/

/*从尾到头打印字符串
 *描述：
 *	输入一个链表，从尾到头打印链表每个节点的值。
 *分析：
 *	常规方法：遍历一遍链表改变next指针的指向。
 *	栈方法：存储到栈中，遍历完后输出。
 *          使用递归函数去遍历链表。（函数的调用方式基于栈）
 */

package main

import (
	"fmt"
)

type ListNode struct {
	value int
	next  *ListNode
}

// 添加元素节点
func (l *ListNode) append(e int) *ListNode {
	node := ListNode{e, nil}
	ptr := l
	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = &node
	return l
}

// 非递归方式
func (l *ListNode) PrintListReverse() {
	// 先存储在，然后再反序列打印
	var slice []int
	p := l
	for p != nil {
		//println(l.value)
		slice = append(slice, p.value)
		p = p.next
	}
	//println(slice)
	for i := len(slice) - 1; i >= 0; i-- {
		fmt.Print(slice[i])
	}
}

// 递归方式
func PrintListReverse1(ptr *ListNode) {
	if ptr != nil {
		if ptr.next != nil {
			PrintListReverse1(ptr.next)
		}
		fmt.Print(ptr.value)
	}
}

func main() {
	var list ListNode
	for i := 0; i < 8; i++ {
		list.append(i)
	}
	fmt.Println("非递遍历：")
	list.PrintListReverse()
	fmt.Println()
	fmt.Println("递归遍历：")
	PrintListReverse1(&list)
}
