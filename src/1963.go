package leetcode_1963

/*
You are given a 0-indexed string s of even length n. The string consists of exactly n / 2 opening brackets '[' and n / 2 closing brackets ']'.
You may swap the brackets at any two indices any number of times.
Return the minimum number of swaps to make s balanced.
*/
func minSwaps(s string) int {
	balance, minUnmatchedBrackets := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			balance++
		} else {
			balance--
		}

		if balance < minUnmatchedBrackets {
			minUnmatchedBrackets = balance
		}
	}

	return (-minUnmatchedBrackets + 1) / 2
}
