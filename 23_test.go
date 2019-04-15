package leetcode

import "testing"

func TestMergeKLists(t *testing.T)  {
	mergeKLists([]*ListNode{
		tranIntListToNode([]int{1,4,5}),
		tranIntListToNode([]int{1,3,4}),
	})
}
