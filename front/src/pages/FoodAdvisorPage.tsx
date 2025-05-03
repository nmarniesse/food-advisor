import { useState } from "react";
import {
  Backdrop,
  Button,
  CircularProgress,
  Container,
  Divider,
  Stack,
  Typography,
} from "@mui/material";
import { DefaultMenuData, MenuData } from "../models/MenuData";
import { useFoodAdvisorForm } from "../hooks/useFoodAdvisorForm";
import { FoodAdvisorResult } from "../components/FoodAdvisorResult";
import { FoodAdvisorForm } from "../components/FoodAdvisorForm";
import { Section } from "../components/common/Section";

const FoodAdvisorPage = () => {
  const [formData, setFormData] = useState<MenuData>(DefaultMenuData);
  const { submit, data: responseData, isLoading } = useFoodAdvisorForm();

  return (
    <Container maxWidth="md">
      {isLoading && (
        <Backdrop
          sx={(theme) => ({ color: "#fff", zIndex: theme.zIndex.drawer + 1 })}
          open={true}
        >
          <CircularProgress color="inherit" />
        </Backdrop>
      )}
      <Typography variant="h2" gutterBottom>
        Générateur de menus
      </Typography>
      <FoodAdvisorForm formData={formData} setFormData={setFormData} />

      <Divider />
      <Section>
        {!responseData && (
          <Button
            variant="contained"
            onClick={() => submit(formData)}
            disabled={isLoading}
          >
            Valider
          </Button>
        )}
        {responseData && (
          <Stack
            spacing={2}
            direction="row"
            sx={{ alignItems: "center", mb: 1 }}
          >
            <Button
              variant="contained"
              onClick={() => submit(formData)}
              disabled={isLoading}
            >
              Relancer
            </Button>
            <Button
              variant="contained"
              color="secondary"
              onClick={() => alert("TODO")}
              disabled={isLoading}
            >
              Relancer en gardant la sélection
            </Button>
          </Stack>
        )}
      </Section>

      {!isLoading && responseData && (
        <FoodAdvisorResult
          recipes={responseData.recipes}
          groceryList={responseData.groceryList}
        />
      )}
    </Container>
  );
};

export { FoodAdvisorPage };
