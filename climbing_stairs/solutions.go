package climbing_stairs

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	n1 := 1
	n2 := 2

	for i := 3; i < n+1; i++ {
		swap := n2
		n2 = n1 + n2
		n1 = swap
	}

	return n2
}
