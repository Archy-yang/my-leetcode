package leetcode

// 可以用二分法，进行两两合并，这样合并list的时间复杂度是O(logN)
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	t := lists[0]

	for i := 1; i < len(lists); i++ {
		if t == nil {
			t = lists[i]
			continue
		}
		if lists[i] == nil {
			continue
		}

		c := lists[i]
		if c == nil {
			continue
		}
		if c.Val < t.Val {
			t, c = c, t
		}
		tmp := t
		for {
			if tmp.Next == nil {
				tmp.Next = c
				break
			}
			if c == nil {
				break
			}
			if tmp.Next.Val > c.Val {
				tmp.Next, c.Next, c = c, tmp.Next, c.Next
			}
			tmp = tmp.Next
		}
	}

	return t
}
