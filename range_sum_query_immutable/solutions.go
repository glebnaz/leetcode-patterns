package range_sum_query_immutable

type NumArray struct {
	arr  []int
	sums map[int]int
}

func Constructor(nums []int) NumArray {
	sums := make(map[int]int)

	currentSum := 0

	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]
		sums[i] = currentSum
	}

	return NumArray{
		arr:  nums,
		sums: sums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sums[right]
	} else {
		return this.sums[right] - this.sums[left-1]
	}
}
