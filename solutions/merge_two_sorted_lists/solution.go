package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mrg := &ListNode{Val: 0}
	res := mrg

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			mrg.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			mrg.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		mrg = mrg.Next
	}

	for list1 != nil {
		mrg.Next = &ListNode{Val: list1.Val}
		list1 = list1.Next
		mrg = mrg.Next
	}
	for list2 != nil {
		mrg.Next = &ListNode{Val: list2.Val}
		list2 = list2.Next
		mrg = mrg.Next
	}
	return res.Next
}
