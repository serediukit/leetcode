package leetcode_ans

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	canBeCooked := make(map[string]bool)

	countCooked := 1
	for countCooked > 0 {
		countCooked = 0
		for i := 0; i < len(recipes); i++ {
			if _, cooked := canBeCooked[recipes[i]]; !cooked {
				canCook := true
				for _, ingredient := range ingredients[i] {
					if !slices.Contains(supplies, ingredient) {
						_, ok := canBeCooked[ingredient]
						if !ok {
							canCook = false
							break
						}
					}
				}
				if canCook {
					canBeCooked[recipes[i]] = true
					countCooked++
				}
			}
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