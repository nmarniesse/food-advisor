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
import { useGetWeekMenus } from "../hooks/useGetWeekMenus";
import { FoodAdvisorResult } from "../components/FoodAdvisorResult";
import { FoodAdvisorForm } from "../components/FoodAdvisorForm";
import { Section } from "../components/common/Section";
import { useRefineWeekMenus } from "../hooks/useRefineWeekMenus";

const FoodAdvisorPage = () => {
  const [formData, setFormData] = useState<MenuData>(DefaultMenuData);
  const { submit, data: responseData, isLoading } = useGetWeekMenus();
  const {
    submit: submitRefine,
    data: responseDataRefine,
    isLoading: isLoadingRefine,
  } = useRefineWeekMenus();

  return (
    <Container maxWidth="md">
      {(isLoading || isLoadingRefine) && (
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
        {(responseData || responseDataRefine) && (
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
              onClick={() => submitRefine(["lundi", "mardi"])}
              disabled={isLoadingRefine}
            >
              Relancer en gardant la sélection
            </Button>
          </Stack>
        )}
      </Section>

      {!isLoading && responseData && !responseDataRefine && (
        <FoodAdvisorResult
          recipes={responseData.recipes}
          groceryList={responseData.groceryList}
        />
      )}
      {!isLoadingRefine && responseDataRefine && (
        <FoodAdvisorResult
          recipes={responseDataRefine.recipes}
          groceryList={responseDataRefine.groceryList}
        />
      )}
    </Container>
  );
};

export { FoodAdvisorPage };
