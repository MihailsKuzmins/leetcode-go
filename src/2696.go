package leetcode_2696

/*
You are given a string s consisting only of uppercase English letters.
You can apply some operations to this string where, in one operation, you can remove any occurrence of one of the substrings "AB" or "CD" from s.
Return the minimum possible length of the resulting string that you can obtain.
Note that the string concatenates after removing the substring and could produce new "AB" or "CD" substrings.
*/
func minLength(s string) int {
	result, j := make([]byte, len(s)), -1

	for i := 0; i < len(s); i++ {
		if s[i] == 'B' {
			if j != -1 && result[j] == 'A' {
				goto Result
			}
		} else if s[i] == 'D' {
			if j != -1 && result[j] == 'C' {
				goto Result
			}
		}

		j++
		result[j] = s[i]
		continue

	Result:
		result[j] = 0
		j--
	}

	return j + 1
}
