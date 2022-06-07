package path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var stack []*TreeNode
	isInter := make(map[*TreeNode]struct{})
	lvl := 0

	var pathSum = 0
	for root != nil || len(stack) > 0 {
		isInter[root] = struct{}{}
		pathSum += root.Val

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

		if (pathSum == targetSum) && (root.Left == nil && root.Right == nil) {
			return true
		}

		if len(stack)-2 < 0 {
			break
		}

		pathSum -= root.Val
		root = stack[len(stack)-2]
		stack = stack[0 : len(stack)-2]
		lvl--
		pathSum -= root.Val
	}
	return false
}
