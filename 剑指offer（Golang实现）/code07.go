/*
 Author: skyler
 Date: 2018-02-02
 */

/*两个栈实现队列
 *描述：
 *  用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。
 *分析：
 *  栈的特点是只可以在栈顶做操作（出栈和入栈），把两个栈分为队列进入栈s1和队列出栈s2。
 *  入队时，将元素压入s1。
 *  出队时，判断s2是否为空，如不为空，则直接弹出顶元素；
 *  如为空，则将s1的元素逐个“倒入”s2，把s2栈顶元素弹出并出队。
 */

package main

import (
	"stack"
	"fmt"
	"errors"
)

type queue struct {
	s1 stack.Stack
	s2 stack.Stack
}

func (q *queue) Pop() interface{} {
	// 判断s2是否为空，如不为空，则直接弹出顶元素；如为空，则将s1的元素逐个“倒入”s2，把s2栈顶元素弹出并出队。
	if q.s2.IsEmpty() {
		size := q.s1.Size()
		for i := 0; i < size; i++ {
			x, _ := q.s1.Pop()
			q.s2.Push(x)
		}
	}
	res, _ := q.s2.Pop()

	// 空接口类型，可用断言的方式获取当前类型输出
	switch res.(type) {
	case int:
		return res.(int)
	case string:
		return res.(string)
	case nil:
		res = errors.New("queue is empty")
	default:
		res = errors.New("type is error")
	}
	return res
}

func (q *queue) Push(e interface{}) {
	// 入队时，将元素压入s1。
	q.s1.Push(e)
}

func main() {
	var q queue
	q.Push(1)
	//q.Push(2)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
