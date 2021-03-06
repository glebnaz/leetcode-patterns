package find_all_numbers_disappeared_in_an_array

//Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.
//
//
//
//Example 1:
//
//Input: nums = [4,3,2,7,8,2,3,1]
//Output: [5,6]
//Example 2:
//
//Input: nums = [1,1]
//Output: [2]
//
//
//Constraints:
//
//n == nums.length
//1 <= n <= 105
//1 <= nums[i] <= n
//
//
//Follow up: Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

func findDisappearedNumbers(nums []int) []int {
	data := make(map[int]bool)

	for i := range nums {
		data[i+1] = false
	}

	for i := range nums {
		data[nums[i]] = true
	}

	var res []int

	for i := range nums {
		if !data[i+1] {
			res = append(res, i+1)
		}
	}

	return res
}

func findDisappearedNumbersBestTry(nums []int) []int {
	i := 0
	for i < len(nums) {
		pos := nums[i] - 1
		if nums[i] != nums[pos] {
			swap := nums[i]

			nums[i] = nums[pos]
			nums[pos] = swap
		} else {
			i++
		}
	}

	var res []int

	for i := range nums {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}

	return res
}
