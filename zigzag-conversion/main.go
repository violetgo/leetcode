package main

import "fmt"

/***
https://leetcode-cn.com/problems/zigzag-conversion/

将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G
 */

func convert(s string, numRows int) string {
	length := len(s)
	if numRows >= length || numRows < 2 {
		return s
	}

	ans:=make([]byte,length)
	j:=0
	for row := 0; row < numRows; row++ {
		i := row
		if row == 0 || row == numRows-1 {
			for ; i < length; {
				ans[j] = s[i]
				j++
				i += numRows*2 - 2
			}
		} else {
			for ; i < length; {
				ans[j] = s[i]
				j++
				i += (numRows - row - 1) * 2
				if i < length {
					ans[j] = s[i]
					j++
					i += row * 2
				}
			}
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(convert("LEETCODEISHIRING", 3))
	fmt.Println(convert("LEETCODEISHIRING", 4))
}
