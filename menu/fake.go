package menu

import (
	"encoding/json"
)

type Fake struct{}

func (fake *Fake) RunQuery(query *Query) (*Response, error) {
	var res Response
	if err := json.Unmarshal(fakeData, &res); err != nil {
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
		{"name": "poulet", "quantity": 500},
		{"name": "salade verte", "quantity": 1},
		{"name": "tomate", "quantity": 2},
		{"name": "vinaigrette", "quantity": 1}
	  ],
	  "preparation": [
		"Faire griller le poulet",
		"Couper la salade, les tomates et le poulet",
		"Ajouter la vinaigrette"
	  ]
	},
	{
	  "day": "Mardi",
	  "recipeName": "Pâtes au jambon et tomates cerises",
	  "ingredients": [
		{"name": "pâtes", "quantity": 300},
		{"name": "jambon", "quantity": 200},
		{"name": "tomates cerises", "quantity": 200},
		{"name": "fromage râpé", "quantity": 100}
	  ],
	  "preparation": [
		"Cuire les pâtes",
		"Faire revenir le jambon et les tomates cerises dans une poêle",
		"Mélanger les pâtes cuites avec le jambon et les tomates",
		"Saupoudrer de fromage râpé"
	  ]
	},
	{
	  "day": "Mercredi",
	  "recipeName": "Saumon en papillote",
	  "ingredients": [
		{"name": "saumon", "quantity": 500},
		{"name": "poireaux", "quantity": 3},
		{"name": "citron", "quantity": 1}
	  ],
	  "preparation": [
		"Préchauffer le four",
		"Couper les poireaux en rondelles",
		"Déposer le saumon et les poireaux sur une feuille de papier sulfurisé",
		"Arroser de jus de citron",
		"Fermer la papillote et enfourner"
	  ]
	},
	{
	  "day": "Jeudi",
	  "recipeName": "Riz sauté au poulet",
	  "ingredients": [
		{"name": "riz", "quantity": 300},
		{"name": "poulet", "quantity": 300},
		{"name": "poivrons rouges", "quantity": 2},
		{"name": "oignon", "quantity": 1}
	  ],
	  "preparation": [
		"Cuire le riz",
		"Faire revenir le poulet, les poivrons et l'oignon dans une poêle",
		"Ajouter le riz cuit et mélanger"
	  ]
	},
	{
	  "day": "Vendredi",
	  "recipeName": "Pizza maison",
	  "ingredients": [
		{"name": "pâte à pizza", "quantity": 1},
		{"name": "jambon", "quantity": 200},
		{"name": "tomate", "quantity": 2},
		{"name": "fromage", "quantity": 200}
	  ],
	  "preparation": [
		"Étaler la pâte à pizza",
		"Garnir de jambon, tomates et fromage",
		"Enfourner jusqu'à ce que le fromage soit fondu et la pâte dorée"
	  ]
	}
  ],
  "groceryList": [
    {"name": "jambon", "quantity": 200},
	{"name": "fromage", "quantity": 200}
  ]
}`)
