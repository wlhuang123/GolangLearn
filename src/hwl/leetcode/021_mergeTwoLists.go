package leetcode

import (
	"hwl/tool/logs"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode .
type ListNode struct {
	Val  int
	Next *ListNode
}

// Test021 .
func Test021() {
	Check021(mergeTwoLists(createList(1, 3), createList(2)), 1, 2, 3)
	Check021(mergeTwoLists(createList(1, 2), createList(3)), 1, 2, 3)
	Check021(mergeTwoLists(createList(), createList()))
	Check021(mergeTwoLists(createList(1), createList()), 1)
	Check021(mergeTwoLists(createList(), createList(1)), 1)
	Check021(mergeTwoLists(createList(1, 5, 9), createList(1, 5, 6, 9)), 1, 1, 5, 5, 6, 9, 9)
}

func createList(v ...int) *ListNode {
	list := new(ListNode)
	rearList := list
	for _, value := range v {
		tmp := ListNode{Val: value}
		rearList.Next = &tmp
		rearList = rearList.Next
	}
	list = list.Next
	return list
}

// Check021 .
func Check021(result *ListNode, expectResult ...int) {
	isok := checkResult(result, expectResult...)
	if isok {
		logs.Println("test ok ")
	} else {
		logs.Println("test failed ", expectResult)
	}
}

func checkResult(result *ListNode, expectResult ...int) bool {
	defer func() {
		if err := recover(); err != nil {
			logs.Println(err)
		}
	}()

	for _, v := range expectResult {
		if v == result.Val {
			result = result.Next
			continue
		}
		return false
	}
	return true
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	result := new(ListNode)
	p := result
	for {
		if l1 == nil {
			p.Next = l2
			return result.Next
		}

		if l2 == nil {
			p.Next = l1
			return result.Next
		}

		if l1.Val < l2.Val {
			p.Next = l1
			p = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			p = l2
			l2 = l2.Next
		}
	}
}
