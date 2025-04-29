import { useState } from "react";
import {
  Box,
  Button,
  Container,
  Divider,
  Slider,
  Stack,
  Switch,
  TextField,
  Typography,
} from "@mui/material";
import AddIcon from "@mui/icons-material/Add";
import RemoveIcon from "@mui/icons-material/Remove";

const defaultMaxPreparationTimeInMin = 45;
const defaultPersons = 4;

const Section = ({ children }) => {
  return (
    <Box component="section" sx={{ margin: "15px" }}>
      {children}
    </Box>
  );
};

const FoodAdvisorForm = () => {
  const [foodInFridgeList, setFoodInFridgeList] = useState<string[]>([""]);
  const [maxPreparationTimeInMin, setMaxPreparationTimeInMin] =
    useState<number>(defaultMaxPreparationTimeInMin);
  const [useSeasonIngredients, setUseSeasonIngredients] =
    useState<boolean>(true);
  const [persons, setPersons] = useState<number>(defaultPersons);

  const updateFoodInFridge = (index: number, value: string) => {
    setFoodInFridgeList([
      ...foodInFridgeList.slice(0, index),
      value,
      ...foodInFridgeList.slice(index + 1),
    ]);
  };
  const removeFoodInFridge = (index: number) => {
    console.log("index", index);

    setFoodInFridgeList([
      ...foodInFridgeList.slice(0, index),
      ...foodInFridgeList.slice(index + 1),
    ]);
  };
  const addFoodInFridge = () => setFoodInFridgeList([...foodInFridgeList, ""]);

  return (
    <Container maxWidth="md">
      <Typography variant="h2" gutterBottom>
        Générateur de menus
      </Typography>
      <Section>
        <Typography component="div">Aliments dans le frigo</Typography>
        {foodInFridgeList.map((foodInFridge, index) => (
          <div key={index}>
            <Stack
              spacing={2}
              direction="row"
              sx={{ alignItems: "center", mb: 1 }}
            >
              <TextField
                label="Ingredient"
                variant="outlined"
                value={foodInFridge}
                onChange={(event) =>
                  updateFoodInFridge(index, event.currentTarget.value)
                }
                margin="normal"
              />
              {index !== 0 && (
                <Button
                  variant="text"
                  onClick={() => removeFoodInFridge(index)}
                >
                  <RemoveIcon />
                </Button>
              )}
              {index === foodInFridgeList.length - 1 && (
                <Button variant="text" onClick={addFoodInFridge}>
                  <AddIcon />
                </Button>
              )}
            </Stack>
          </div>
        ))}
      </Section>
      <Divider />
      <Section>
        <Typography component="div">temps de préparation maximum</Typography>
        <Stack spacing={4} direction="row" sx={{ alignItems: "center", mb: 1 }}>
          <Slider
            size="medium"
            defaultValue={defaultMaxPreparationTimeInMin}
            aria-label="Max preparation time in minutes"
            valueLabelDisplay="auto"
            step={15}
            min={15}
            max={90}
            onChange={(_, newValue) => setMaxPreparationTimeInMin(newValue)}
            sx={{ width: "250px" }}
          />
          <div>{maxPreparationTimeInMin} minutes</div>
        </Stack>
      </Section>
      <Divider />
      <Section>
        <Typography component="div">
          Utiliser des ingrédients de saison?
        </Typography>
        <Switch
          checked={useSeasonIngredients}
          onClick={() => setUseSeasonIngredients(!useSeasonIngredients)}
        />
      </Section>
      <Divider />
      <Section>
        <Typography component="div">Personnes</Typography>
        <Stack spacing={4} direction="row" sx={{ alignItems: "center", mb: 1 }}>
          <Slider
            size="medium"
            defaultValue={defaultPersons}
            aria-label="Persons"
            valueLabelDisplay="auto"
            step={1}
            min={1}
            max={10}
            onChange={(_, newValue) => setPersons(newValue)}
            sx={{ width: "250px" }}
          />
          <div>{persons}</div>
        </Stack>
      </Section>
      <Divider />
      <Section>
        <Button variant="contained" onClick={() => alert("todo")}>
          Valider
        </Button>
      </Section>
    </Container>
  );
};

export { FoodAdvisorForm };
