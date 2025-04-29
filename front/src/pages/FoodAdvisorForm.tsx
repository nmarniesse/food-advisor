import { FC, useState } from "react";
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
import { DefaultMenuData, MenuData } from "../models/MenuData";
import { useFoodAdvisorForm } from "../hooks/useFoodAdvisorForm";

const Section: FC<{ children: ReactDOM }> = ({ children }) => {
  return (
    <Box component="section" sx={{ margin: "15px" }}>
      {children}
    </Box>
  );
};

const FoodAdvisorForm = () => {
  const [data, setData] = useState<MenuData>(DefaultMenuData);
  const { submit } = useFoodAdvisorForm();

  const updateFoodInFridge = (index: number, value: string) => {
    setData({
      ...data,
      foodInFridgeList: [
        ...data.foodInFridgeList.slice(0, index),
        value,
        ...data.foodInFridgeList.slice(index + 1),
      ],
    });
  };
  const removeFoodInFridge = (index: number) => {
    console.log("index", index);

    setData({
      ...data,
      foodInFridgeList: [
        ...data.foodInFridgeList.slice(0, index),
        ...data.foodInFridgeList.slice(index + 1),
      ],
    });
  };
  const addFoodInFridge = () =>
    setData({ ...data, foodInFridgeList: [...data.foodInFridgeList, ""] });

  return (
    <Container maxWidth="md">
      <Typography variant="h2" gutterBottom>
        Générateur de menus
      </Typography>
      <Section>
        <Typography component="div">Aliments dans le frigo</Typography>
        {data.foodInFridgeList.map((foodInFridge, index) => (
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
              {index === data.foodInFridgeList.length - 1 && (
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
            value={data.maxPreparationTimeInMin}
            aria-label="Max preparation time in minutes"
            valueLabelDisplay="auto"
            step={15}
            min={15}
            max={90}
            onChange={(_, newValue) =>
              setData({ ...data, maxPreparationTimeInMin: newValue })
            }
            sx={{ width: "250px" }}
          />
          <div>{data.maxPreparationTimeInMin} minutes</div>
        </Stack>
      </Section>
      <Divider />
      <Section>
        <Typography component="div">Personnes</Typography>
        <Stack spacing={4} direction="row" sx={{ alignItems: "center", mb: 1 }}>
          <Slider
            size="medium"
            value={data.persons}
            aria-label="Persons"
            valueLabelDisplay="auto"
            step={1}
            min={1}
            max={10}
            onChange={(_, newValue) => setData({ ...data, persons: newValue })}
            sx={{ width: "250px" }}
          />
          <div>{data.persons}</div>
        </Stack>
      </Section>
      <Divider />
      <Section>
        <Typography component="div">
          Utiliser des ingrédients de saison?
        </Typography>
        <Switch
          checked={data.useSeasonIngredient}
          onClick={() =>
            setData({ ...data, useSeasonIngredient: !data.useSeasonIngredient })
          }
        />
      </Section>
      <Divider />
      <Section>
        <Button variant="contained" onClick={() => submit(data)}>
          Valider
        </Button>
      </Section>
    </Container>
  );
};

export { FoodAdvisorForm };
