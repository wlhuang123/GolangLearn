package leetcode

/*
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/

// Test203 .
func Test203() {

}

// Check203 .
func Check203() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	result := new(ListNode)
	p := result
	p.Next = head
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return result.Next
}

// 方法二：删除头节点特殊考虑
func removeElements2(head *ListNode, val int) *ListNode {
	for {
		if head == nil {
			return nil
		}

		if head.Val == val {
			head = head.Next
			continue
		}
		break
	}

	p := head
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return head
}
