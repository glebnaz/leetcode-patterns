package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	middle := slow

	var prev *ListNode = nil

	current := middle
	var next *ListNode

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	for prev != nil {
		if prev.Val != head.Val {
			return false
		}
		prev = prev.Next
		head = head.Next
	}

	return true
}
