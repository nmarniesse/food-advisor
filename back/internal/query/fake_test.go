package query

import (
	"testing"

	"github.com/nmarniesse/food-advisor/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestItReturnsFakeData(t *testing.T) {
	f := &Fake{}
	query := &model.Query{[]string{}, 30, true, 3}

	res, err := f.RunQuery(query)
	assert.Nil(t, err)

	assert.Len(t, res.Menus, 5)
	assert.Len(t, res.GroceryList, 2)
	assert.Equal(t, "Lundi", res.Menus[0].Day)
	assert.Equal(t, "Salade de poulet grillé", res.Menus[0].RecipeName)
	assert.Equal(t, "poulet", res.Menus[0].Ingredients[0].Name)
	assert.Equal(t, "500g", res.Menus[0].Ingredients[0].Quantity)
	assert.Equal(t, "Mardi", res.Menus[1].Day)
}
