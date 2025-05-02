export type MenuData = {
  foodInFridgeList: string[];
  maxPreparationTimeInMin: number;
  useSeasonIngredient: boolean;
  persons: number;
};

export const DefaultMenuData: MenuData = {
  foodInFridgeList: [""],
  maxPreparationTimeInMin: 45,
  useSeasonIngredient: true,
  persons: 4,
};

export type Recipe = {
  day: string;
  recipeName: string;
  ingredients: Ingredient[];
  preparation: string[];
  recipeLink: string;
};

export type Ingredient = {
  name: string;
  quantity: number;
};
