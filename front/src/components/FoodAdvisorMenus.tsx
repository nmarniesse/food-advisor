import {
  Box,
  Collapse,
  Grid,
  Paper,
  Tab,
  Tabs,
  Typography,
} from "@mui/material";
import { Ingredient, Recipe } from "../models/MenuData";
import { Section } from "./common/Section";
import { useState } from "react";
import { ExpandLess, ExpandMore } from "@mui/icons-material";

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
  const [openPreparations, setOpenPreparations] = useState<boolean[]>(
    recipes.map(() => false),
  );

  const openPreparation = (index: number) => {
    setOpenPreparations((prev) => {
      const newOpenPreparations = [...prev];
      newOpenPreparations[index] = !newOpenPreparations[index];

      return newOpenPreparations;
    });
  };

  return (
    <Section>
      <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
        <Tabs value={tab} onChange={(_, newTab) => setTab(newTab)}>
          <Tab label="Menu de la semaine" />
          <Tab label="Liste de courses" />
        </Tabs>
      </Box>
      <CustomTabPanel value={tab} index={0}>
        {recipes.length > 0 && (
          <Grid container spacing={2}>
            {recipes.map((recipe: Recipe, index: number) => (
              <Grid size={6} key={index}>
                <Paper>
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
                  <p onClick={() => openPreparation(index)}>
                    Préparation
                    {openPreparations[index] ? <ExpandLess /> : <ExpandMore />}
                  </p>
                  <Collapse
                    in={openPreparations[index] ?? false}
                    timeout="auto"
                    unmountOnExit
                  >
                    <ol>
                      {recipe.preparation.map((step, index) => (
                        <li key={index}>{step}</li>
                      ))}
                    </ol>
                  </Collapse>
                </Paper>
              </Grid>
            ))}
          </Grid>
        )}
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
