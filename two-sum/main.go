package main

import "fmt"

/***
https://leetcode-cn.com/problems/two-sum/submissions/

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func twoSum(nums []int, target int) []int {
	index := []int{-1, -1}
	sumSet := make(map[int][]int)
	for i, k := range nums {
		if sumSet[k] == nil {
			t := make([]int, 0)
			sumSet[k] = t
		}
		sumSet[k] = append(sumSet[k], i)
	}
	for i, k := range sumSet {
		if len(k) > 1 && i+i == target {
			return k
		}
		if target-i == i {
			continue
		}
		num := sumSet[target-i]
		if num != nil {
			return []int{k[0], num[0]}
		}
	}

	return index
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{3, 0, 1, 2, 3}, 6))
	fmt.Println(twoSum([]int{1,3,4,2}, 6))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
