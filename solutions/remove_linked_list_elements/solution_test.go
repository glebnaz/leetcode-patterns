package remove_linked_list_elements

import (
	"testing"
)

func Test_removeElements(t *testing.T) {

	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	removeElements(head, 2)
}
