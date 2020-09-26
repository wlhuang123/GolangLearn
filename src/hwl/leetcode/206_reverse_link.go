package leetcode

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/

// Test206 .
func Test206() {

}

// Check206 .
func Check206() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	q := head.Next
	p.Next = nil

	for q.Next != nil {
		temp := q.Next

		q.Next = p
		p = q
		q = temp
	}
	q.Next = p
	return q
}
