package leetcode_ans

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	recipeToIngredients := make(map[string][]string)
	visited := make(map[string]int)
	res := make([]string, 0, len(recipes))

	for i, recipe := range recipes {
		recipeToIngredients[recipe] = ingredients[i]
	}

	var canCooked func(string) bool
	canCooked = func(recipe string) bool {
		if val, exists := visited[recipe]; exists {
			return val == 1
		}

		if slices.Contains(supplies, recipe) {
			return true
		}

		if _, exists := recipeToIngredients[recipe]; !exists {
			return false
		}

		visited[recipe] = 0

		for _, ingredient := range recipeToIngredients[recipe] {
			if !canCooked(ingredient) {
				visited[recipe] = -1
				return false
			}
		}

		visited[recipe] = 1
		res = append(res, recipe)
		return true
	}

	for _, recipe := range recipes {
		canCooked(recipe)
	}

	return res
}