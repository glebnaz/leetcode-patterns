package palindrome_linked_list

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
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

	if !isPalindrome(head) {
		fmt.Printf("is not true")
		t.FailNow()
	}
}
