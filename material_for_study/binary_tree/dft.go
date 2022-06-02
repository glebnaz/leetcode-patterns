package binary_tree

import "fmt"

func dftGlebNaz(root *TreeNode) {
	if root == nil {
		return
	}

	var stack []*TreeNode
	isInter := make(map[*TreeNode]struct{})
	lvl := 0

	for root != nil || len(stack) > 0 {
		isInter[root] = struct{}{}

		//действие
		fmt.Printf("lvl: %v val: %v\n", lvl, root.Val)

		stack = append(stack, root)

		if root.Left != nil {
			if _, ok := isInter[root.Left]; !ok {
				root = root.Left
				lvl++
				continue
			}
		}

		if root.Right != nil {
			if _, ok := isInter[root.Right]; !ok {
				root = root.Right
				lvl++
				continue
			}
		}

		if len(stack)-2 < 0 {
			break
		}

		root = stack[len(stack)-2]
		stack = stack[0 : len(stack)-2]
		lvl--
	}
}
