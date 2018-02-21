/*
 *Author: skyler
 *Date: 2018-02-16
 */

/*斐波那契额数列
 *  构造斐波那契数列，要求输入一个整数n，请你输出斐波那契数列的第n项。
 *
 *跳台阶
 *描述：
 *  一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
 *分析：
 *  当要跳n次时用f(n)表示，最后一跳有两种方式跳1级f(n-1)和跳2级f(n-2)，f(n)=f(n-1)+f(n-2)，即斐波那契数列。
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
	fibs[n] = res
	return res
}

func Fibonacci(n int) uint {
	i, a, b := 0, 0, 1
	for i < n {
		a, b = b, a+b
		i++
	}
	return uint(a)
}

func main() {
	fmt.Println(Fibonacci(10))
	fmt.Println(FibonacciRecursive(10))
}
