package leetcode_3

/*
Given a string s, find the length of the longest substring without repeating characters.
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLength, currentLength := 0, 0
	indexMap := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		if startIndex, ok := indexMap[s[i]]; ok {
			if currentLength > maxLength {
				maxLength = currentLength
			}

			currentLength = i - startIndex

			for key, value := range indexMap {
				if value < startIndex {
					delete(indexMap, key)
				}
			}
		} else {
			currentLength++
		}

		indexMap[s[i]] = i
	}

	if maxLength > currentLength {
		return maxLength
	} else {
		return currentLength
	}
}
