/*
 *Author: skyler
 *Date: 2018-02-16
 */

/*斐波那契额数列
 *  构造斐波那契数列，要求输入一个整数n，请你输出斐波那契数列的第n项。
 */

package main

import "fmt"

const LIM = 100

var fibs [LIM]uint

func FibonacciRecursive(n int) (res uint) {
	if fibs[n] != 0 { // 构造缓存避免递归时重复计算
		res = fibs[n]
		return
	}

	if n <= 1 {
		return uint(n)
	}
	res = FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
	return res
}

func Fibonacci(n int) uint {
	i, a, b := 0, 0, 1
	for i < n {
		a, b = b, a+b
		i ++
	}
	return uint(a)
}

func main() {
	fmt.Println(Fibonacci(10))
	fmt.Println(FibonacciRecursive(10))
}
