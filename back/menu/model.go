package menu

type Query struct {
	FoodInFridge            []string
	MaxPreparationTimeInMin int
	UseSeasonIngredient     bool
	Persons                 int
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

type Menu struct {
	Day         string       `json:"day"`
	RecipeName  string       `json:"recipeName"`
	Ingredients []Ingredient `json:"ingredients"`
	Preparation []string     `json:"preparation"`
	RecipeLink  string       `json:"recipeLink"`
}

type Response struct {
	Menus       []Menu       `json:"recipes"`
	GroceryList []Ingredient `json:"groceryList"`
}

type RefineQuery struct {
	DaysToKeep []string
}

type RefineResponse struct {
	Response string
}

type IAProvider interface {
	RunQuery(Query) (*Response, error)
	RunRefineQuery(RefineQuery) (*Response, error)
}
