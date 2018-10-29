package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	t := &ListNode{}
	for head != nil {
		t.Val = head.Val
		tmp := &ListNode{}
		tmp.Next = t
		t = tmp
		head = head.Next
	}

	return t.Next
}

func reverseList1 (head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	t, _ := reversRe(head)

	return t
}

func reversRe (head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}
	t, d := reversRe(head.Next)
	d.Next = &ListNode{
		Val:head.Val,
	}
	d = d.Next

	return t, d
}
