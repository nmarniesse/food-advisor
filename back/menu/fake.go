package menu

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
)

type Fake struct{}

func (f *Fake) RunQuery(query *Query) (*Response, error) {
	message := query.formatToString()
	log.Println("Query is about to start with message:\n" + message)

	time.Sleep(1 * time.Second)
	var res Response
	if err := json.Unmarshal(fakeData, &res); err != nil {
		return nil, err
	}

	res.Uuid = uuid.New()

	return &res, nil
}

func (f *Fake) RunRefineQuery(query *RefineQuery) (*Response, error) {
	time.Sleep(1 * time.Second)
	var res Response
	if err := json.Unmarshal(fakeDataRefine, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

var fakeData = []byte(`{
  "recipes": [
	{
	  "day": "Lundi",
	  "recipeName": "Salade de poulet grillé",
	  "ingredients": [
		{"name": "poulet", "quantity": "500g"},
		{"name": "salade verte", "quantity": "1"},
		{"name": "tomate", "quantity": "2"},
		{"name": "vinaigrette", "quantity": "1"}
	  ],
	  "preparation": [
		"Faire griller le poulet",
		"Couper la salade, les tomates et le poulet",
		"Ajouter la vinaigrette"
	  ],
	  "recipeLink": "https://www.cuisineaz.com/recettes/salade-de-poulet-grille-67481.aspx"
	},
	{
	  "day": "Mardi",
	  "recipeName": "Pâtes au jambon et tomates cerises",
	  "ingredients": [
		{"name": "pâtes", "quantity": "300g"},
		{"name": "jambon", "quantity": "200g"},
		{"name": "tomates cerises", "quantity": "200g"},
		{"name": "fromage râpé", "quantity": "100g"}
	  ],
	  "preparation": [
		"Cuire les pâtes",
		"Faire revenir le jambon et les tomates cerises dans une poêle",
		"Mélanger les pâtes cuites avec le jambon et les tomates",
		"Saupoudrer de fromage râpé"
	  ],
	  "recipeLink": "https://www.cuisineaz.com/recettes/salade-de-poulet-grille-67481.aspx"
	},
	{
	  "day": "Mercredi",
	  "recipeName": "Saumon en papillote",
	  "ingredients": [
		{"name": "saumon", "quantity": "500g"},
		{"name": "poireaux", "quantity": "3"},
		{"name": "citron", "quantity": "1"}
	  ],
	  "preparation": [
		"Préchauffer le four",
		"Couper les poireaux en rondelles",
		"Déposer le saumon et les poireaux sur une feuille de papier sulfurisé",
		"Arroser de jus de citron",
		"Fermer la papillote et enfourner"
	  ],
	  "recipeLink": "https://www.cuisineaz.com/recettes/salade-de-poulet-grille-67481.aspx"
	},
	{
	  "day": "Jeudi",
	  "recipeName": "Riz sauté au poulet",
	  "ingredients": [
		{"name": "riz", "quantity": "300g"},
		{"name": "poulet", "quantity": "300g"},
		{"name": "poivrons rouges", "quantity": "2"},
		{"name": "oignon", "quantity": "1"}
	  ],
	  "preparation": [
		"Cuire le riz",
		"Faire revenir le poulet, les poivrons et l'oignon dans une poêle",
		"Ajouter le riz cuit et mélanger"
	  ],
	  "recipeLink": "https://www.cuisineaz.com/recettes/salade-de-poulet-grille-67481.aspx"
	},
	{
	  "day": "Vendredi",
	  "recipeName": "Pizza maison",
	  "ingredients": [
		{"name": "pâte à pizza", "quantity": "1"},
		{"name": "jambon", "quantity": "200g"},
		{"name": "tomate", "quantity": "2"},
		{"name": "fromage", "quantity": "200g"}
	  ],
	  "preparation": [
		"Étaler la pâte à pizza",
		"Garnir de jambon, tomates et fromage",
		"Enfourner jusqu'à ce que le fromage soit fondu et la pâte dorée"
	  ],
	  "recipeLink": "https://www.cuisineaz.com/recettes/salade-de-poulet-grille-67481.aspx"
	}
  ],
  "groceryList": [
    {"name": "jambon", "quantity": "200g"},
	{"name": "fromage", "quantity": "200g"}
  ]
}`)

var fakeDataRefine = []byte(`{
	"recipes": [
	  {
		"day": "Lundi",
		"recipeName": "Salade de poulet grillé [Refine]",
		"ingredients": [
		  {"name": "poulet", "quantity": "500g"},
		  {"name": "salade verte", "quantity": "1"},
		  {"name": "tomate", "quantity": "2"},
		  {"name": "vinaigrette", "quantity": "1"}
		],
		"preparation": [
		  "Faire griller le poulet",
		  "Couper la salade, les tomates et le poulet",
		  "Ajouter la vinaigrette"
		],
		"recipeLink": "https://www.cuisineaz.com/recettes/salade-de-poulet-grille-67481.aspx"
	  },
	  {
		"day": "Mardi",
		"recipeName": "Pâtes au jambon et tomates cerises [Refine]",
		"ingredients": [
		  {"name": "pâtes", "quantity": "300g"},
		  {"name": "jambon", "quantity": "200g"},
		  {"name": "tomates cerises", "quantity": "200g"},
		  {"name": "fromage râpé", "quantity": "100g"}
		],
		"preparation": [
		  "Cuire les pâtes",
		  "Faire revenir le jambon et les tomates cerises dans une poêle",
		  "Mélanger les pâtes cuites avec le jambon et les tomates",
		  "Saupoudrer de fromage râpé"
		],
		"recipeLink": "https://www.cuisineaz.com/recettes/salade-de-poulet-grille-67481.aspx"
	  },
	  {
		"day": "Mercredi",
		"recipeName": "Saumon en papillote [Refine]",
		"ingredients": [
		  {"name": "saumon", "quantity": "500g"},
		  {"name": "poireaux", "quantity": "3"},
		  {"name": "citron", "quantity": "1"}
		],
		"preparation": [
		  "Préchauffer le four",
		  "Couper les poireaux en rondelles",
		  "Déposer le saumon et les poireaux sur une feuille de papier sulfurisé",
		  "Arroser de jus de citron",
		  "Fermer la papillote et enfourner"
		],
		"recipeLink": "https://www.cuisineaz.com/recettes/salade-de-poulet-grille-67481.aspx"
	  },
	  {
		"day": "Jeudi",
		"recipeName": "Riz sauté au poulet [Refine]",
		"ingredients": [
		  {"name": "riz", "quantity": "300g"},
		  {"name": "poulet", "quantity": "300g"},
		  {"name": "poivrons rouges", "quantity": "2"},
		  {"name": "oignon", "quantity": "1"}
		],
		"preparation": [
		  "Cuire le riz",
		  "Faire revenir le poulet, les poivrons et l'oignon dans une poêle",
		  "Ajouter le riz cuit et mélanger"
		],
		"recipeLink": "https://www.cuisineaz.com/recettes/salade-de-poulet-grille-67481.aspx"
	  },
	  {
		"day": "Vendredi",
		"recipeName": "Pizza maison [Refine]",
		"ingredients": [
		  {"name": "pâte à pizza", "quantity": "1"},
		  {"name": "jambon", "quantity": "200g"},
		  {"name": "tomate", "quantity": "2"},
		  {"name": "fromage", "quantity": "200g"}
		],
		"preparation": [
		  "Étaler la pâte à pizza",
		  "Garnir de jambon, tomates et fromage",
		  "Enfourner jusqu'à ce que le fromage soit fondu et la pâte dorée"
		],
		"recipeLink": "https://www.cuisineaz.com/recettes/salade-de-poulet-grille-67481.aspx"
	  }
	],
	"groceryList": [
	  {"name": "jambon", "quantity": "200g"},
	  {"name": "fromage", "quantity": "200g"}
	]
  }`)
