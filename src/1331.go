package leetcode_1331

import "sort"

/*
Given an array of integers arr, replace each element with its rank.
The rank represents how large the element is. The rank has the following rules:
*/
func arrayRankTransform(arr []int) []int {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	ranksMap := map[int]int{}
	for i, rank := 0, 1; i < len(sortedArr); i = i + 1 {
		if _, exists := ranksMap[sortedArr[i]]; !exists {
			ranksMap[sortedArr[i]] = rank
			rank++
		}
	}

	ranks := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		ranks[i] = ranksMap[arr[i]]
	}

	return ranks
}
