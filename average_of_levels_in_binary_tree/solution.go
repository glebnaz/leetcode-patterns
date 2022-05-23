package average_of_levels_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64

	if root == nil {
		return res
	}

	lvlNodes := []*TreeNode{root}

	for len(lvlNodes) != 0 {
		sum := 0
		count := 0

		nextLvl := make([]*TreeNode, 0, 0)
		for _, node := range lvlNodes {
			sum = sum + node.Val
			count++

			if node.Left != nil {
				nextLvl = append(nextLvl, node.Left)
			}

			if node.Right != nil {
				nextLvl = append(nextLvl, node.Right)
			}
		}

		res = append(res, float64(sum)/float64(count))
		lvlNodes = nextLvl
	}
	return res
}
