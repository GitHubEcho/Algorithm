/*
 Author: skyler
 Date: 2018-01-31
*/

package main

import "fmt"

/* 冒泡排序
 时间复杂度 O(n^2)
 稳定
 优化：有序序列重复多次-->在某次冒泡后元素没有交换位置，直接返回
*/
func BubbleSort(sli []int) {
	for i := len(sli) - 1; i > 0; i-- {
		isChange := false
		for j := 0; j < i; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
				isChange = true
			}
			if !isChange {
				return
			}
		}
	}
}

/* 选择排序
 时间复杂度 O(n^2)
 稳定
*/
func SelectSort(sli []int) {
	for i := 0; i < len(sli); i++ {
		min := i
		for j := i + 1; j < len(sli); j++ {
			if sli[min] > sli[j] {
				min = j
			}
		}
		sli[i], sli[min] = sli[min], sli[i]
	}

}

/*堆排序

 */

/* 插入排序
 时间复杂度O(n^2)
 稳定
*/
func InsertSort(sli []int) {
	for i := 1; i < len(sli); i++ {
		for j := i; j > 0; j-- {
			if sli[j] < sli[j-1] {
				sli[j], sli[j-1] = sli[j-1], sli[j]
			}
		}
	}
}


/* 希尔排序：在插入排序的基础上，加入了gap分而治之。
 * 时间复杂度O(n^2)
 * 稳定
 * 优化：所选的gap刚好是有序的且不互为质数，算法退化为插入排序-->增量序列互为质数。
 */
func ShellSort(sli []int) {
	n := len(sli)
	for D := n / 2; D > 0; D /= 2 {
		for i := D; i < n; i++ {
			for j := i; j >= D ; j -= D {
				if sli[j] < sli[j-D] {
					sli[j], sli[j-D] = sli[j-D], sli[j]
				}
			}
		}
	}
}

func QuickSort() {

}

func main() {
	sli := []int{7, 4, 9, 3, 10, 5}
	fmt.Println(sli)
	//BubbleSort(sli)
	//SelectSort(sli)


	//InsertSort(sli)
	ShellSort(sli)
	fmt.Print(sli)
}
