package find_smallest_letter_greater_than_target

func nextGreatestLetter(letters []byte, target byte) byte {
	start := 0
	end := len(letters) - 1

	for end >= start {
		mid := start + (end-start)/2
		if letters[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return letters[start%len(letters)]
}
