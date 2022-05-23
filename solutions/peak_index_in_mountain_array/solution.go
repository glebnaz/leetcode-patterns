package peak_index_in_mountain_array

func peakIndexInMountainArray(nums []int) int {
	start := 0
	end := len(nums) - 1

	var mid int

	for start <= end {
		//сдвиг плюс расстояние между началом и концом, деленное на два
		mid = start + (end-start)/2

		if nums[mid] < nums[mid+1] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}
