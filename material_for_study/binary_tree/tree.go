package binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//type stack struct {
//	data *TreeNode
//	next *stack
//}
//
//func (s *stack) Add(d *TreeNode) *stack {
//	return &stack{data: d, next: s}
//}
//
//func (s *stack) Remove() *TreeNode {
//	if s.next == nil {
//		return nil
//	}
//	d := s.next.data
//	s.next = s.next.next
//	return d
//}
