package maximum_subarray

// решение с "разделяй и властвуй я так и не смог реализовать"

//https://www.geeksforgeeks.org/maximum-subarray-sum-using-divide-and-conquer-algorithm/ тут подсказка

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
