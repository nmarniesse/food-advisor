import { FC, useState } from "react";
import {
  Box,
  Button,
  Card,
  CardActionArea,
  CardActions,
  CardContent,
  Checkbox,
  Collapse,
  Grid,
  Link,
  List,
  ListItem,
  ListItemText,
  Stack,
  Typography,
} from "@mui/material";
import { Recipe } from "../models/MenuData";

type Props = {
  recipe: Recipe;
  isSelected: boolean;
  onToggleSelection: () => void;
};

const MenuView: FC<Props> = ({ recipe, isSelected, onToggleSelection }) => {
  const [isPreparationOpen, setIsPreparationOpen] = useState<boolean>(false);

  const togglePreparationView = () => {
    setIsPreparationOpen((prev) => !prev);
  };

  return (
    <Card variant="outlined">
      <CardActionArea onClick={onToggleSelection}>
        <CardContent>
          <Stack
            direction={"row"}
            spacing={2}
            sx={{ mb: 2, alignItems: "center" }}
          >
            <Typography
              gutterBottom
              sx={{ color: "text.secondary", fontSize: 14 }}
            >
              {recipe.day}
            </Typography>
            <Checkbox checked={isSelected} />
          </Stack>
          <Typography variant="h5">{recipe.recipeName}</Typography>
          <Box sx={{ marginTop: "10px" }}>Ingrédients</Box>
          <List dense={true}>
            {recipe.ingredients.map((ingredient, index) => (
              <ListItem key={index}>
                <ListItemText
                  primary={`${ingredient.name} - ${ingredient.quantity}`}
                />
              </ListItem>
            ))}
          </List>
          <Link href={recipe.recipeLink} target="_blank">
            Voir la rectte complète
          </Link>
        </CardContent>
      </CardActionArea>
      <CardActions>
        <Button onClick={togglePreparationView}>
          {isPreparationOpen ? "Cacher Préparation" : "Voir Préparation"}
        </Button>
      </CardActions>
      <Collapse in={isPreparationOpen} timeout="auto" unmountOnExit>
        <CardContent>
          <ol>
            {recipe.preparation.map((step, index) => (
              <li key={index}>{step}</li>
            ))}
          </ol>
        </CardContent>
      </Collapse>
    </Card>
  );
};

export { MenuView };
