import { useState } from "react";
import { Button, Container, TextField } from "@mui/material";
import AddIcon from "@mui/icons-material/Add";

const FoodAdvisorForm = () => {
  const [foodInFridgeList, setFoodInFridgeList] = useState<string[]>([""]);

  const updateFoodInFridge = (index: number, value: string) => {
    setFoodInFridgeList([
      ...foodInFridgeList.slice(0, index),
      value,
      ...foodInFridgeList.slice(index + 1),
    ]);
  };
  const addFoodInFridge = () => setFoodInFridgeList([...foodInFridgeList, ""]);

  return (
    <Container maxWidth="md">
      {foodInFridgeList.map((foodInFridge, index) => (
        <div>
          <TextField
            label="Aliment"
            variant="outlined"
            value={foodInFridge}
            onChange={(event) =>
              updateFoodInFridge(index, event.currentTarget.value)
            }
            margin="normal"
          />
          {index === foodInFridgeList.length - 1 && (
            <Button variant="text" onClick={addFoodInFridge}>
              <AddIcon />
            </Button>
          )}
        </div>
      ))}
    </Container>
  );
};

export { FoodAdvisorForm };
