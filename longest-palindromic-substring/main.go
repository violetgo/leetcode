package main

import "fmt"

/**
https://leetcode-cn.com/problems/longest-palindromic-substring/

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func longestPalindrome(s string) string {
	var left, right, maxLength, startIndex, index int
	len := len(s)

	for ; index < len; {
		right = index
		left = index - 1
		for {
			if right >= len || s[right] != s[index] {
				break
			}
			right++
		}

		//定位下一个子串的中心
		index = right

		//以连续重复字符为中心，往左右延展,判断当前子串是否为回文子串
		for {
			if left < 0 || right >= len || s[left] != s[right] {
				break
			}
			left--
			right++
		}
		//记录回文子串的最大长度和起始索引
		if right-left-1 > maxLength {
			startIndex = left + 1
			maxLength = right - left - 1
		}
	}

	return s[startIndex:startIndex+maxLength]
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome(""))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("bb"))
	fmt.Println(longestPalindrome("12345678900987654321"))
	fmt.Println(longestPalindrome("123456789000987654321"))

}
