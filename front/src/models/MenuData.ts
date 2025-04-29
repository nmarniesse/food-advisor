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
