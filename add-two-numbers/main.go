package main

import "fmt"

/**

https://leetcode-cn.com/problems/add-two-numbers/

给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getNextOrEnd(l *ListNode) *ListNode {
	if l == nil {
		return l
	} else {
		return l.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		l1 = &ListNode{}
	}

	if l2 != nil {
		l1.Val += l2.Val
	}

	if l1.Val > 9 {
		l1.Val = l1.Val % 10

		if l1.Next == nil {
			l1.Next = &ListNode{Val: 1}
		} else {
			l1.Next.Val += 1
		}

	}
	l1.Next = addTwoNumbers(l1.Next, getNextOrEnd(l2))

	return l1
}

func printStruct(l *ListNode) {
	if l == nil {
		fmt.Println()
		return
	}
	printStruct(l.Next)
	fmt.Print(l.Val)
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	printStruct(addTwoNumbers(l1, l2))

	l1 = &ListNode{Val: 5}
	l2 = &ListNode{Val: 5}
	printStruct(addTwoNumbers(l1, l2))

	l1 = &ListNode{Val: 1}
	l2 = &ListNode{Val: 9, Next: &ListNode{Val: 9}}
	printStruct(addTwoNumbers(l1, l2))
}
