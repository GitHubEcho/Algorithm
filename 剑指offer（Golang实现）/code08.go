/*
 *Author: skyler
 *Date: 2018-02-05
 */

/*旋转数组的最小数字
 *描述：
 *  把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
 *  输入一个非递减排序的数组的一个旋转，输出旋转数组的最小元素。
 *  例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。
 *  NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。
 *分析：
 *  常规的解法为重头到尾进行一次冒泡排序，找出最小的即可，但是此种方法时间复杂度为O(n)。
 *  显然不能充分利用旋转数组的特性，可以采用二分查找的方式充分使用数组的有序性。
 */

package main

import (
	"errors"
	"fmt"
)

func Min(numbers []int) int {
	length := len(numbers)
	if length <= 0 || numbers == nil {
		errors.New("Invaild parameters !")
	}
	index1 := 0
	index2 := length - 1
	indexmid := index1 // 当数组第一个数小于最后一个数时，数组是有序数列，可以直接返回。
	for numbers[index1] >= numbers[index2] {
		if index2-index1 == 1 {
			indexmid = index2
			break
		}

		indexmid = (index1 + index2 ) / 2

		// 当出现三个数相等时只能使用顺序的方式查找
		if numbers[indexmid] == numbers[index1] && numbers[index1] == numbers[index2] {
			return minOrder(numbers)
		}

		if numbers[indexmid] >= numbers[index1] {
			index1 = indexmid
		} else if numbers[indexmid] <= numbers[index2] {
			index2 = indexmid
		}
	}
	return numbers[indexmid]
}

func minOrder(numbers []int) int {
	min := numbers[0]
	for _, v := range numbers {
		if min > v {
			min = v
		}
	}
	return min
}

func main() {
	numbers1 := []int{3, 4, 5, 1, 2}
	numbers2 := []int{1, 0, 1, 1, 1}
	fmt.Println(Min(numbers1))
	fmt.Println(Min(numbers2))
}
