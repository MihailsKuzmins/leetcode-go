package leetcode_1497

/*
Given an array of integers arr of even length n and an integer k.
We want to divide the array into exactly n / 2 pairs such that the sum of each pair is divisible by k.
Return `true` If you can find a way to do that or `false` otherwise.
*/
func canArrange(arr []int, k int) bool {
	frequency := make([]int, k)

	for i := 0; i < len(arr); i++ {
		remainder := ((arr[i] % k) + k) % k
		frequency[remainder]++
	}

	if frequency[0]%2 != 0 {
		return false
	}

	for i, maxLen := 1, k/2; i <= maxLen; i++ {
		if frequency[i] != frequency[k-i] {
			return false
		}
	}

	return true
}
