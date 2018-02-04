/*
 Author: skyler
 Date: 2018-02-02
 */

/*两个栈实现队列
 *描述：
 *	用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。
 *分析：
 *  入队时，将元素压入s1。
 *  出队时，判断s2是否为空，如不为空，则直接弹出顶元素；
 *  如为空，则将s1的元素逐个“倒入”s2，把最后一个元素弹出并出队。
 */

package main

import (
	"stack"
)

//
//
//

var stack stack.Stack



