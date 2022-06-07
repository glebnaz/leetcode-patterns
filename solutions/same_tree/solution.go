package same_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil || q == nil {
		return false
	}

	lvlNodesp := []*TreeNode{p}
	lvlNodesq := []*TreeNode{q}

	for len(lvlNodesp) != 0 && len(lvlNodesq) != 0 {

		nextLvlp := make([]*TreeNode, 0, 0)
		nextLvlq := make([]*TreeNode, 0, 0)

		if len(lvlNodesp) != len(lvlNodesq) {
			return false
		}

		//p
		for i, node := range lvlNodesp {

			if node.Left != nil {
				nextLvlp = append(nextLvlp, node.Left)
			}

			if node.Right != nil {
				nextLvlp = append(nextLvlp, node.Right)
			}
			if lvlNodesq[i].Val != node.Val {
				return false
			}
		}

		//q
		for _, node := range lvlNodesq {
			if node.Left != nil {
				nextLvlq = append(nextLvlq, node.Left)
			}

			if node.Right != nil {
				nextLvlq = append(nextLvlq, node.Right)
			}
		}

		lvlNodesp = nextLvlp
		lvlNodesq = nextLvlq
	}
	return true
}
