package comparer

import (
	reader "Go_Day01/cmd/readDB/reader"
	"fmt"
)

func CompareCakes(oldRecipe, newRecipe *reader.Recipe) {
	oldCakes := make(map[string]reader.Cake)
	newCakes := make(map[string]reader.Cake)

	for _, cake := range oldRecipe.Cakes {
		oldCakes[cake.Name] = cake
	}
	for _, cake := range newRecipe.Cakes {
		newCakes[cake.Name] = cake
	}

	for name, oldCake := range oldCakes {
		if newCake, exists := newCakes[name]; exists {
			CompareCakeDetails(name, oldCake, newCake)
			delete(newCakes, name)
		} else {
			fmt.Printf("REMOVED cake \"%s\"\n", name)
		}
	}

	for name := range newCakes {
		fmt.Printf("ADDED cake \"%s\"\n", name)
	}
}

func CompareCakeDetails(name string, oldCake, newCake reader.Cake) {
	if oldCake.Time != newCake.Time {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", name, newCake.Time, oldCake.Time)
	}

	oldIngredients := make(map[string]reader.Ingredient)
	newIngredients := make(map[string]reader.Ingredient)

	for _, ing := range oldCake.Ingredients {
		oldIngredients[ing.Name] = ing
	}

	for _, ing := range newCake.Ingredients {
		newIngredients[ing.Name] = ing
	}

	for ingName, oldIng := range oldIngredients {
		if newIng, exists := newIngredients[ingName]; exists {
			CompareIngredientDetails(name, oldIng, newIng)
			delete(newIngredients, ingName)
		} else {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", ingName, name)
		}
	}

	for ingName := range newIngredients {
		fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", ingName, name)
	}
}

func CompareIngredientDetails(cakeName string, oldIng, newIng reader.Ingredient) {
	if oldIng.Unit != newIng.Unit {
		fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldIng.Name, cakeName, newIng.Unit, oldIng.Unit)
	}

	if oldIng.Count != newIng.Count {
		fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldIng.Name, cakeName, newIng.Count, oldIng.Count)
	}
}
