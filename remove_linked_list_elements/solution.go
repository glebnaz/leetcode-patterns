package remove_linked_list_elements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	//случайно начал решать не ту задачу
	//data := make(map[int]*ListNode)
	//
	//curr := head
	//var prev *ListNode = nil
	//
	//for curr != nil {
	//	node, ok := data[curr.Val]
	//	if ok {
	//		if node == head {
	//			node = node.Next
	//		}
	//		node.Next = node.Next.Next
	//		prev.Next = curr.Next
	//		curr = curr.Next
	//		continue
	//	}
	//	curr = curr.Next
	//	prev = curr
	//
	//	data[curr.Val] = prev
	//}
	//return head

	prev := &ListNode{Val: 0, Next: head}

	curr := prev

	for curr.Next != nil {

		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return prev.Next
}
