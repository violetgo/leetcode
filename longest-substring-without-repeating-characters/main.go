package main

import (
	"fmt"
)

/***
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func lengthOfLongestSubstring(s string) int {
	var size, charIndex, leftIndex, innerIndex, max int
	size = len(s)
	if size == 0 {
		return 0
	} else {
		max = 1
	}
	for charIndex = 0; charIndex < size; charIndex++ {
		for innerIndex = leftIndex; innerIndex < charIndex; innerIndex++ {
			fmt.Println(max,innerIndex,charIndex,s[innerIndex:charIndex], s[innerIndex] == s[charIndex])
			if s[innerIndex] == s[charIndex] {
				leftIndex = innerIndex + 1
				break
			}
			if charIndex-leftIndex > max {
				max = charIndex - leftIndex
			}
		}
	}

	if size-leftIndex > max {
		return size - leftIndex
	} else {
		return max
	}
}

func main() {
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("  "))
	fmt.Println(lengthOfLongestSubstring("au"))
	fmt.Println(lengthOfLongestSubstring("cdd"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
