package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func min(num1 int, num2 int) int {
	if num1 > num2 {
		return num2

	} else {
		return num1
	}
}

func findKth(nums1 []int, i int, nums2 []int, j int, k int) int {
	l1 := len(nums1)
	if i >= l1 {
		return nums2[j+k-1] //nums1为空数组
	}
	l2 := len(nums2)
	if j >= l2 {
		return nums1[i+k-1] //nums2为空数组
	}
	if k == 1 {
		return min(nums1[i], nums2[j])
	}

	var midVal1, midVal2 int
	i0 := i + k/2
	i1 := i0 - 1
	if i1 < l1 {
		midVal1 = nums1[i1]
	} else {
		midVal1 = math.MaxInt32
	}

	j0 := j + k/2
	j1 := j0 - 1
	if j1 < l2 {
		midVal2 = nums2[j1]
	} else {
		midVal2 = math.MaxInt32
	}

	if midVal1 < midVal2 {
		return findKth(nums1, i0, nums2, j, k-k/2)
	} else {
		return findKth(nums1, i, nums2, j0, k-k/2)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	s := m + n + 1
	left := s / 2
	right := (s + 1) / 2

	n1 := findKth(nums1, 0, nums2, 0, left)
	n2 := findKth(nums1, 0, nums2, 0, right)

	return float64(n1+n2) / 2
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1}, []int{2, 3, 4, 5, 6}))
	//3.5
}
