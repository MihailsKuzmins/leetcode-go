package leetcode_921

/*
You are given a parentheses string s. In one move, you can insert a parenthesis at any position of the string.
For example, if s = "()))", you can insert an opening parenthesis to be "(()))" or a closing parenthesis to be "())))".
Return the minimum number of moves required to make s valid.
*/
func minAddToMakeValid(s string) int {
	balance, moves := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			balance++
		} else {
			balance--
		}

		if balance < 0 {
			moves++
			balance = 0
		}
	}

	return balance + moves
}
