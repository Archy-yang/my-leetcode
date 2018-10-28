package leetcode


func tranIntListToNode(ints []int) *ListNode {
	var tn = &ListNode{}
	node := tn
	for _, i := range ints {
		t := &ListNode{
			i,
			nil,
		}
		node.Next = t
		node = node.Next
	}

	return tn.Next
}

func transListNodeToIntList(node *ListNode) []int {
	ints := []int{}

	for node != nil {
		ints = append(ints, node.Val)
		node = node.Next
	}

	return ints
}
