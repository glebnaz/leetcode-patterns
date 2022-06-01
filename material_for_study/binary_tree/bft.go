package binary_tree

import "fmt"

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
	var lvl = 0

	fmt.Printf("lvl: 0 Root %v\n", root.Val)

	for len(lvlNodes) != 0 {
		lvl++
		sum := 0
		count := 0

		nextLvl := make([]*TreeNode, 0, 0)
		for _, node := range lvlNodes {
			sum = sum + node.Val
			count++

			if node.Left != nil {
				fmt.Printf("lvl: %v left %v\n", lvl, node.Left.Val)
				nextLvl = append(nextLvl, node.Left)
			}

			if node.Right != nil {
				fmt.Printf("lvl: %v right %v\n", lvl, node.Right.Val)
				nextLvl = append(nextLvl, node.Right)
			}
		}

		res = append(res, float64(sum)/float64(count))
		lvlNodes = nextLvl
	}
	return res
}
