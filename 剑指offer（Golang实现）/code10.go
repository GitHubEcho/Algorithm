/*
 *Author: skyler
 *Date: 2018-02-17
 */

/*二进制中1的个数
 *描述：
 *  输入一个整数，输出该数二进制表示中1的个数。其中负数用补码表示。
 *相关：
 *  所用知识点：移位运算和与运算。这些运算往往会给我们带来便利性。
 *  &运算后返回值为0或1，移位运算3（00000011）右移为1（00000001）。
 *分析：
 *  ①每次原数n向右移动然后做与运算，返回值为int类型1和0，当为1时count计数加1
 *  ②在①的基础上，避免输入n的数字为负数时，发生错误。用左移1与原数字做与运算。
 *  ③此种方法充分利用二进制的特性，把数n减去1和n做与运算，会把n的最右边少一个1变为0，循环做此种运算直到n的值为0。
 */
package main

import "fmt"

func NumberOfOne1(n int) int {
	count := 0
	for n != 0 { // bool(n)编译错误，go语言中1不表示ture。
		if (n & 1) != 0 {
			count ++
		}
		n = n >> 1
	}
	return count
}

func NumberOfOne2(n int) int {
	flag := 1
	count := 0
	for flag != 0 {
		if (n & flag) != 0 {
			count ++
		}
		flag = flag << 1
	}
	return count
}

func NumberOfOne3(n int) int {
	count := 0
	for n != 0 {
		count ++
		n = (n - 1) & n

	}
	return count
}


func main() {
	fmt.Println(NumberOfOne1(3))
	fmt.Println(NumberOfOne2(3))
	fmt.Println(NumberOfOne3(3))
}