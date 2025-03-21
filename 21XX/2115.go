package leetcode_ans

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	canBeCooked := make(map[string]bool)

	var checkCooking func(string, []string)
	checkCooking = func(recipe string, ingredients []string) {
		canCook := true
		for _, ingredient := range ingredients {
			if !slices.Contains(supplies, ingredient) {
				_, ok := canBeCooked[ingredient]
				if !ok {
					canCook = false
					break
				}
			}
		}
		if canCook {
			canBeCooked[recipe] = true
		}
	}

	for i := 0; i < len(recipes); i++ {
		if _, cooked := canBeCooked[recipes[i]]; !cooked {
			checkCooking(recipes[i], ingredients[i])
		}
	}

	res := make([]string, len(canBeCooked))
	index := 0
	for recipe := range canBeCooked {
		res[index] = recipe
		index++
	}

	return res
}