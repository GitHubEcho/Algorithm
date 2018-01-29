/*
 Author: skyler
 Date: 2018-01-28
*/

/*替换空格
  描述：
     请实现一个函数，将一个字符串中的空格替换成“%20”。例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy
  分析：
	常规的方法是：把字符串遍历，遇到空格替换字符串。这样会导致最后的字符串，会随空格个数增加。
	最优解：先计算替换后字符串的长度。从字符串的后面遍历替换字符串。
*/

package main

import (
	"fmt"
	//"strings"
)

// str：传入的字符串
// rep：替换的字符串
func ReplaceBlack(str string, rep string) string {
	str_arr := make([]rune, 20) // 字符串为非ASCII（如汉字）
	//str_arr := []byte(str)    // 字符串为ASCII
	rep_arr := []rune(rep)

	if str_arr != nil {
		//统计空白个数
		countBlank := 0
		for i, v := range str {
			str_arr[i] = v
			if string(v) == " " {
				countBlank ++
			}
		}

		// 计算新字符串的长度、原字符和新字符串的索引
		strLengthNew := len(str) + countBlank*(len(rep)-1)
		indexOfOld := len(str) - 1
		indexOfNew := strLengthNew - 1

		for indexOfOld >= 0 && indexOfNew > indexOfOld {
			if string(str_arr[indexOfOld]) == " " {
				for j := len(rep_arr); j > 0; j-- {
					str_arr[indexOfNew] = rep_arr[j-1]
					indexOfNew --
				}
			} else {
				str_arr[indexOfNew] = str_arr[indexOfOld]
				indexOfNew --
			}
			indexOfOld --
		}
	}
	return string(str_arr)
}

func main() {
	str := "We Are Happy"
	//fmt.Println(strings.Replace(str," ","%20",-1)) //string自带的字符串替换
	fmt.Println(ReplaceBlack(str, "%20"))
}
