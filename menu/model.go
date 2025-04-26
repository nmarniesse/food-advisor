package menu

import (
	"fmt"
	"strings"
)

type Query struct {
	FoodInFridge            []string
	MaxPreparationTimeInMin int
	UseSeasonIngredient     bool
	Persons                 int
}

type IAProvider interface {
	RunQuery(Query) (*Response, error)
}

func (query *Query) formatToString() string {
	template := `
Je souhaite que tu me fasses un menu pour la semaine avec les conditions suivantes :
Ingrédients disponibles dans mon frigo : %s. Tu peux en rajouter si besoin. Essaie de ne pas utiliser le même ingrédient plus de 1 repas.
Temps de préparation : Chaque repas doit être prêt en %d minutes maximum.
Nombre de personnes : %d personnes. Donne moi juste le plat principal.
%s
Recettes et liste de courses : Pour chaque repas, donne-moi la recette et la liste des courses globale pour toute la semaine.
Je vais interpreter le résultat donc donne moi le resultat en format json comme par exemple:
{
  "recipes": [
    {
      "day": "lundi",
      "recipeName": "Omelette",
      "ingredients": [
      	{"name": "oeufs", "quantity": 3},
		{"name": "jambon", "quantity": 200}
      ],
      "preparation": [
      	"battre les oeufs",
      	"cuire"
      ]
    }
  ],
  "groceryList": [
    {"name": "oeufs", "quantity": 3},
	{"name": "jambon", "quantity": 200}
  ]
]
`
	var seasonIngredientSentence string
	if query.UseSeasonIngredient {
		seasonIngredientSentence = `Produits de saison : Utilise uniquement des produits de saison, car on est au printemps.`
	} else {
		seasonIngredientSentence = ""
	}

	return fmt.Sprintf(
		template,
		strings.Join(query.FoodInFridge, ", "),
		query.MaxPreparationTimeInMin, query.Persons,
		seasonIngredientSentence,
	)
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Menu struct {
	Day         string       `json:"day"`
	RecipeName  string       `json:"recipeName"`
	Ingredients []Ingredient `json:"ingredients"`
	Preparation []string     `json:"preparation"`
}

type Response struct {
	Menus       []Menu       `json:"recipes"`
	GroceryList []Ingredient `json:"groceryList"`
}
