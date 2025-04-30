import { useState } from "react";
import { Ingredient, MenuData, Recipe } from "../models/MenuData";

type Return = {
  submit: (menuData: MenuData) => void;
  data: MenuResponse | null;
  isLoading: boolean;
  error: string | null;
};

type MenuResponse = {
  recipes: Recipe[];
  groceryList: Ingredient[];
};

const useFoodAdvisorForm = (): Return => {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [data, setData] = useState<MenuResponse | null>(null);

  const submit = (menuData: MenuData) => {
    console.log(menuData);
    console.log("VITE_BACKEND_URL", import.meta.env.VITE_BACKEND_URL);

    setIsLoading(true);
    setError(null);
    setData(null);
    fetch(`${import.meta.env.VITE_BACKEND_URL}/week-menu`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then(async (response) => {
        setIsLoading(false);
        if (!response.ok) {
          console.error("Error response:", response);
          setError(response.statusText);

          return;
        }

        setData(await response.json());
      })
      .catch((error) => {
        setIsLoading(false);
        console.error("Error fetching data:", error);
        setError(error.message);
      });
  };

  return { submit, data, isLoading, error };
};

export { useFoodAdvisorForm };
