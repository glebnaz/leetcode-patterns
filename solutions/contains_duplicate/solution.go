package containsduplicate

//Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
//
//
//
//Example 1:
//
//Input: nums = [1,2,3,1]
//Output: true
//Example 2:
//
//Input: nums = [1,2,3,4]
//Output: false
//Example 3:
//
//Input: nums = [1,1,1,3,3,4,3,2,4,2]
//Output: true

//https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
	data := make(map[int]int)

	for i := range nums {
		data[nums[i]] = data[nums[i]] + 1
		if data[nums[i]] > 1 {
			return true
		}
	}
	return false
}
