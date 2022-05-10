package maximum_subarray

func maxSubArray(nums []int) int {
	maxSubSum := nums[0]

	current := nums[0]

	max := func(i, g int) int {
		if i > g {
			return i
		}
		return g
	}

	for i := 1; i < len(nums); i++ {

		current = max(current+nums[i], nums[i])

		maxSubSum = max(current, maxSubSum)

	}

	return maxSubSum
}
