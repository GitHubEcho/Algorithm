/*
 *Author: skyler
 *Date: 2018-02-16
 */

/*跳台阶
 *描述：
 *  一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
 *分析：
 *  当要跳n次时用f(n)表示，最后一跳有两种方式跳1级f(n-1)和跳2级f(n-2)，f(n)=f(n-1)+f(n-2)，即斐波那契数列。
 */

package main

import "fmt"

const LIM = 100

var fibs [LIM]uint

func Fib(n int) (res uint) {
	if fibs[n] != 0 {
		res = fibs[n]
	}
	if n <= 1 {
		return uint(n)
	} else {
		res = Fib(n-1) + Fib(n-2)
	}
	fibs[n] = res
	return
}

func main() {
	fmt.Println(Fib(10))
}
