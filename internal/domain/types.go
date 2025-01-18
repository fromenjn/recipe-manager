package domain

type RecipeIllustration struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Filepath    string `json:"filepath"`
}

type RecipeStep struct {
	ID                 string               `json:"id"`
	Name               string               `json:"name"`
	Instructions       string               `json:"instructions"`
	RecipeIllustration []RecipeIllustration `json:"illustration"`
}

type Ingredient struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type Recipe struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
	Steps       []RecipeStep `json:"steps"`
}
