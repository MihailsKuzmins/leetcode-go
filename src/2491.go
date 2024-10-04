package leetcode_2491

/*
You are given a positive integer array skill of even length n where skill[i] denotes the skill of the ith player. Divide the players into n / 2 teams of size 2 such that the total skill of each team is equal.
The chemistry of a team is equal to the product of the skills of the players on that team.
Return the sum of the chemistry of all the teams, or return -1 if there is no way to divide the players into teams such that the total skill of each team is equal.
*/
func dividePlayers(skill []int) int64 {
	if len(skill) < 2 {
		return int64(skill[0] * skill[1])
	}

	var skillSum int64 = 0

	for i := 0; i < len(skill); i++ {
		skillSum += int64(skill[i])
	}

	var chemistry int64 = 0
	teamScore := int(skillSum / int64(len(skill)/2))
	scoreMap := make(map[int]int)

	for i := 0; i < len(skill); i++ {
		search := teamScore - skill[i]
		if count, ok := scoreMap[search]; ok {
			chemistry += int64(search * skill[i])

			if count <= 1 {
				delete(scoreMap, search)
			} else {
				scoreMap[search]--
			}
		} else {
			scoreMap[skill[i]]++
		}
	}

	if len(scoreMap) == 0 {
		return chemistry
	} else {
		return -1
	}
}
