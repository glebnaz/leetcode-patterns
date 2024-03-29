package diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 0
	}

	left := root.Left
	right := root.Right

	maxLeft := maxDepth(left)
	maxRight := maxDepth(right)

	return maxLeft + maxRight
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := -1

	var stack []*TreeNode
	isInter := make(map[*TreeNode]struct{})
	lvl := 1

	for root != nil || len(stack) > 0 {
		isInter[root] = struct{}{}

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

		if root.Left == nil && root.Right == nil {
			if max < lvl {
				max = lvl
			}
		}

		if len(stack)-2 < 0 {
			break
		}

		root = stack[len(stack)-2]
		stack = stack[0 : len(stack)-2]
		lvl--
	}

	return max
}
