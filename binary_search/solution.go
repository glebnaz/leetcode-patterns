package binary_search

const notFound = -1

func search(nums []int, target int) int {
	return bs(nums, 0, len(nums)-1, target)
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
