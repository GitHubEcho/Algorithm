/*
 Author: skyler
 Date: 2018-1-28
*/

/*二维数组中的查找
 在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
 请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
 <p/>
 规律：首先选取数组中右上角的数字。如果该数字等于要查找的数字，查找过程结束：
 如果该数字大于要查找的数字，剔除这个数字所在的列：如果该数字小于要查找的数字，剔除这个数字所在的行。
 也就是说如果要查找的数字不在数组的右上角，则每－次都在数组的查找范围中剔除）行或者一列，这样每一步都可以缩小
 查找的范围，直到找到要查找的数字，或者查找范围为空。
*/

package main

import "fmt"

// matrix 待查找的数组
// rows columns 行列数
// number 要查找数字
func Find(matrix  [][] int, number int) bool {
	// 获取数组行列数
	rows := len(matrix)
	columns := len(matrix[0])

	// 定位右上角
	row := 0
	column := columns - 1
	found := false

	// 从右上角开始查找，当前值大于查找值时，要找的元素可能在左边（列序号减一）
	// 当小于时在可能在当前数的下边（行序号减一）
	if matrix != nil && rows > 0 && columns > 0{
		for row < rows && column >= 0 {
			if matrix[row][column] == number{
				found = true
				break
			}else if number < matrix[row][column]{
				column --
			}else {
				row ++
			}
		}
	}
	return found
}

func main() {
	matrix := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	//fmt.Println(len(matrix[1]))
	for i :=0;i < 15;i++{
		fmt.Printf("find %d :%v\n",i,Find(matrix,i))
	}
}
