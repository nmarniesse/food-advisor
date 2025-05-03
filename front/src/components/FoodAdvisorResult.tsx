import { FC, useState } from "react";
import {
  Box,
  Grid,
  List,
  ListItem,
  ListItemText,
  Tab,
  Tabs,
} from "@mui/material";
import { Ingredient, Recipe } from "../models/MenuData";
import { Section } from "./common/Section";
import { MenuView } from "./MenuView";

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

const FoodAdvisorResult: FC<Props> = ({ recipes, groceryList }) => {
  const [tab, setTab] = useState<number>(0);
  const [menuSelected, setMenuSelectedsetOpenPreparations] = useState<
    boolean[]
  >(recipes.map(() => false));

  const toggleSelection = (index: number) => {
    setMenuSelectedsetOpenPreparations((prev) => {
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
              <Grid size={6}>
                <MenuView
                  recipe={recipe}
                  isSelected={menuSelected[index] ?? false}
                  onToggleSelection={() => toggleSelection(index)}
                />
              </Grid>
            ))}
          </Grid>
        )}
        {recipes.length === 0 && <p>Aucun menu trouvé.</p>}
      </CustomTabPanel>
      <CustomTabPanel value={tab} index={1}>
        <List dense={true}>
          {groceryList.map((ingredient: Ingredient, index: number) => (
            <ListItem key={index}>
              <ListItemText
                primary={`${ingredient.name} - ${ingredient.quantity}`}
              />
            </ListItem>
          ))}
        </List>
      </CustomTabPanel>
    </Section>
  );
};

export { FoodAdvisorResult };
