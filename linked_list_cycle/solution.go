package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// нужно использовать быстрый и медленный указатель
// подробнее
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head
	slow := head

	for fast.Next != nil && slow.Next != nil {
		if fast.Next.Next == nil {
			return false
		}
		if slow.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
