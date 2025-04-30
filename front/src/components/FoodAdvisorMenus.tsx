import { Box, Tab, Tabs, Typography } from "@mui/material";
import { Ingredient, Recipe } from "../models/MenuData";
import { Section } from "./common/Section";
import { useState } from "react";

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

const CustomTabPanel = (props: TabPanelProps) => {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ p: 3 }}>{children}</Box>}
    </div>
  );
};

type Props = {
  recipes: Recipe[];
  groceryList: Ingredient[];
};

const FoodAdvisorMenus: FC<Props> = ({ recipes, groceryList }) => {
  const [tab, setTab] = useState<number>(0);

  return (
    <Section>
      <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
        <Tabs
          value={tab}
          onChange={(_, newTab) => setTab(newTab)}
          aria-label="basic tabs example"
        >
          <Tab label="Menu de la semaine" />
          <Tab label="Liste de courses" />
        </Tabs>
      </Box>
      <CustomTabPanel value={tab} index={0}>
        {recipes.map((recipe: Recipe, index: number) => (
          <div key={index}>
            <Typography variant="h3">{recipe.day}</Typography>
            <Typography variant="h4">{recipe.recipeName}</Typography>
            <p>Ingrédients :</p>
            <ul>
              {recipe.ingredients.map((ingredient, index) => (
                <li key={index}>
                  {ingredient.name} : {ingredient.quantity}
                </li>
              ))}
            </ul>
            <p>Préparation :</p>
            <ol>
              {recipe.preparation.map((step, index) => (
                <li key={index}>{step}</li>
              ))}
            </ol>
          </div>
        ))}
        {recipes.length === 0 && <p>Aucun menu trouvé.</p>}
      </CustomTabPanel>
      <CustomTabPanel value={tab} index={1}>
        <ul>
          {groceryList.map((ingredient: Ingredient, index: number) => (
            <li key={index}>
              {ingredient.name} : {ingredient.quantity}
            </li>
          ))}
        </ul>
      </CustomTabPanel>
    </Section>
  );
};

export { FoodAdvisorMenus };
