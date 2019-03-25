package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0  {
		return head
	}
	var l int
	t := head
	for t != nil {
		l++
		t = t.Next
	}

	s := head
	for i :=1; i < l-n; i++{
		s = s.Next
	}
	if s == head {
		return s.Next
	}
	s.Next = s.Next.Next

	return head
}