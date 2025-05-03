import { FC } from "react";
import {
  Button,
  Divider,
  Slider,
  Stack,
  Switch,
  TextField,
  Typography,
} from "@mui/material";
import { Section } from "./common/Section";
import AddIcon from "@mui/icons-material/Add";
import RemoveIcon from "@mui/icons-material/Remove";
import { MenuData } from "../models/MenuData";

type Props = {
  formData: MenuData;
  setFormData: (data: MenuData) => void;
};

const FoodAdvisorForm: FC<Props> = ({ formData, setFormData }) => {
  const updateFoodInFridge = (index: number, value: string) => {
    setFormData({
      ...formData,
      foodInFridgeList: [
        ...formData.foodInFridgeList.slice(0, index),
        value,
        ...formData.foodInFridgeList.slice(index + 1),
      ],
    });
  };
  const removeFoodInFridge = (index: number) => {
    setFormData({
      ...formData,
      foodInFridgeList: [
        ...formData.foodInFridgeList.slice(0, index),
        ...formData.foodInFridgeList.slice(index + 1),
      ],
    });
  };
  const addFoodInFridge = () =>
    setFormData({
      ...formData,
      foodInFridgeList: [...formData.foodInFridgeList, ""],
    });

  return (
    <>
      <Section>
        <Typography component="div">Aliments dans le frigo</Typography>
        {formData.foodInFridgeList.map((foodInFridge, index) => (
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
              {index === formData.foodInFridgeList.length - 1 && (
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
        <Typography component="div">Temps de préparation maximum</Typography>
        <Stack spacing={4} direction="row" sx={{ alignItems: "center", mb: 1 }}>
          <Slider
            size="medium"
            value={formData.maxPreparationTimeInMin}
            aria-label="Max preparation time in minutes"
            valueLabelDisplay="auto"
            step={15}
            min={15}
            max={90}
            onChange={(_, newValue) =>
              setFormData({ ...formData, maxPreparationTimeInMin: newValue })
            }
            sx={{ width: "250px" }}
          />
          <div>{formData.maxPreparationTimeInMin} minutes</div>
        </Stack>
      </Section>
      <Divider />
      <Section>
        <Typography component="div">Personnes</Typography>
        <Stack spacing={4} direction="row" sx={{ alignItems: "center", mb: 1 }}>
          <Slider
            size="medium"
            value={formData.persons}
            aria-label="Persons"
            valueLabelDisplay="auto"
            step={1}
            min={1}
            max={10}
            onChange={(_, newValue) =>
              setFormData({ ...formData, persons: newValue })
            }
            sx={{ width: "250px" }}
          />
          <div>{formData.persons}</div>
        </Stack>
      </Section>
      <Divider />
      <Section>
        <Typography component="div">
          Utiliser des ingrédients de saison?
        </Typography>
        <Switch
          checked={formData.useSeasonIngredient}
          onClick={() =>
            setFormData({
              ...formData,
              useSeasonIngredient: !formData.useSeasonIngredient,
            })
          }
        />
      </Section>
    </>
  );
};

export { FoodAdvisorForm };
