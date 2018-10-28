package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	t := &ListNode{}
	cursor1 := t
	cursor2 := head

	for cursor2 != nil {
		if cursor2.Val != val {
			cursor1.Next = &ListNode{
				Val: cursor2.Val,
			}
			cursor1 = cursor1.Next
		}
		cursor2 = cursor2.Next
	}

	return t.Next
}


