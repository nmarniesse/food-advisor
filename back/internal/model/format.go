package model

import (
	"fmt"
	"strings"
)

func (query *Query) FormatToString() string {
	template := `Je souhaite que tu me fasses un menu pour la semaine avec les conditions suivantes :
%s
Essaie de ne pas utiliser le même ingrédient plus de 2 repas.
Temps de préparation : Chaque repas doit être prêt en %d minutes maximum.
Nombre de personnes : %d personnes. Donne moi juste le plat principal.
%s
Les recettes doivent être bien notées sur internet. Tu peux utiliser des sites comme Marmiton, Cuisine AZ ou le journal des femmes.
Recettes et liste de courses : Pour chaque repas, donne-moi la recette et la liste des courses globale pour toute la semaine.
Je vais interpreter le résultat donc donne moi le resultat en format JSON. C'est important que le résultat soit un JSON valide.
Voici un exemple de résultat :
{
  "recipes": [
    {
      "day": "lundi",
      "recipeName": "Omelette",
      "ingredients": [
      	{"name": "oeufs", "quantity": "3"},
        {"name": "jambon", "quantity": "200g"}
      ],
      "preparation": [
      	"battre les oeufs",
      	"cuire"
      ],
      "recipeLink": "https://www.marmiton.org/recettes/recette_omelette-jambon-fromage_68824.aspx"
    }
  ],
  "groceryList": [
    {"name": "oeufs", "quantity": "3"},
    {"name": "jambon", "quantity": "200g"}
  ]
]
`
	var foodInFridgeSentence string
	if len(query.FoodInFridge) > 0 {
		foodInFridgeSentence = fmt.Sprintf("Ingrédients disponibles dans mon frigo : %s. Tu peux en rajouter si besoin. ", strings.Join(query.FoodInFridge, ", "))
	} else {
		foodInFridgeSentence = ""
	}

	var seasonIngredientSentence string
	if query.UseSeasonIngredient {
		seasonIngredientSentence = `Produits de saison : Utilise majoritairement des produits de saison.`
	} else {
		seasonIngredientSentence = ""
	}

	return fmt.Sprintf(
		template,
		foodInFridgeSentence,
		query.MaxPreparationTimeInMin, query.Persons,
		seasonIngredientSentence,
	)
}
