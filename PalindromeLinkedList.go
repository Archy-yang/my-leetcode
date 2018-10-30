package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome1(head *ListNode) bool {
	c := 0
	t := head
	for t != nil {
		c++
		t = t.Next
	}
	if c < 2 {
		return true
	}

	t = head
	head = head.Next
	m := head.Next
	t.Next = nil
	for i:= 1; i< c/2; i++ {
		head.Next = t
		t = head
		head = m
		m = m.Next
	}
	if c % 2 != 0 {
		head = head.Next
	}
	for head != nil {
		if t == nil || t.Val != head.Val {
			return false
		}
		head = head.Next
		t = t.Next
	}
	if t != nil {
		return false
	}

	return true
}
