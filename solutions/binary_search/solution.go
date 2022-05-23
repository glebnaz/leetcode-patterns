package binary_search

const notFound = -1

func search(nums []int, target int) int {
	//by cycle

	start := 0
	end := len(nums) - 1

	var mid int

	for end >= start {
		//сдвиг плюс расстояние между началом и концом, деленное на два
		mid = start + (end-start)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return notFound

	//by recursion
	//return bs(nums, 0, len(nums)-1, target)
}

func bs(arr []int, start, end, target int) int {
	if end >= start {
		mid := start + (end-start)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] > target {
			return bs(arr, start, mid-1, target)
		}

		return bs(arr, mid+1, end, target)
	}

	return notFound
}
