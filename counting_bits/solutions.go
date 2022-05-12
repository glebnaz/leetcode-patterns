package counting_bits

func countBits(n int) []int {
	var data = make([]int, n+1)

	for i := range data {
		data[i] = data[i>>1] + i%2
	}
	return data
}
