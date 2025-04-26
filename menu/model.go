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

type Menu struct {
	day         string
	recipeName  string
	ingredients []string
}

type IAProvider interface {
	RunQuery(Query) []Menu
}

func (query *Query) formatToString() string {
	template := `
Je souhaite que tu me fasses un menu pour la semaine avec les conditions suivantes :
Ingrédients disponibles dans mon frigo : %s. Tu peux en rajouter si besoin. Essaie de ne pas utiliser le même ingrédient plus de 2 repas.
Temps de préparation : Chaque repas doit être prêt en %d minutes maximum.
Nombre de personnes : %d personnes. Donne moi juste le plat principal.
%s
Recettes et liste de courses : Pour chaque repas, donne-moi la recette et la liste des courses globale pour toute la semaine.
Je veux le resultat en format json comme par exemple:
[
  "jour": "lundi",
  "nom_recette": "Omelette",
  "ingredients": [
  	{"nom": "oeufs", "quantite": 3}
  ],
  "preparation": [
  	"battre les oeufs",
  	"cuire"
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
