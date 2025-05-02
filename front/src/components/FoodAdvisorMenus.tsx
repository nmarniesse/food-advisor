import {
  Box,
  Button,
  Card,
  CardContent,
  Collapse,
  Grid,
  List,
  ListItem,
  ListItemText,
  Tab,
  Tabs,
  Typography,
} from "@mui/material";
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
                <Card variant="outlined">
                  <CardContent>
                    <Typography
                      gutterBottom
                      sx={{ color: "text.secondary", fontSize: 14 }}
                    >
                      {recipe.day}
                    </Typography>
                    <Typography variant="h5">{recipe.recipeName}</Typography>
                    <Box sx={{ marginTop: "10px" }}>Ingrédients</Box>
                    <List dense={true}>
                      {recipe.ingredients.map((ingredient, index) => (
                        <ListItem key={index}>
                          <ListItemText
                            primary={`${ingredient.name} - ${ingredient.quantity}`}
                          />
                          {/* {ingredient.name} : {ingredient.quantity} */}
                        </ListItem>
                      ))}
                    </List>
                    <Button onClick={() => openPreparation(index)}>
                      {openPreparations[index]
                        ? "Cacher Préparation"
                        : "Voir Préparation"}
                    </Button>
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
                  </CardContent>
                </Card>
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
