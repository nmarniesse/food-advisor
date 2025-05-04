import { useState } from "react";
import { Ingredient, Recipe } from "../models/MenuData";

type Return = {
  submit: (uuid: string, daysToKeep: string[]) => void;
  data: MenuResponse | null;
  isLoading: boolean;
  error: string | null;
};

export type MenuResponse = {
  uuid: string;
  recipes: Recipe[];
  groceryList: Ingredient[];
};

const useRefineWeekMenus = (): Return => {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [data, setData] = useState<MenuResponse | null>(null);

  const submit = (uuid: string, daysToKeep: string[]) => {
    const params = new URLSearchParams({
      uuid,
      days_to_keep: daysToKeep.join(","),
    });

    setIsLoading(true);
    setError(null);
    setData(null);
    fetch(`${import.meta.env.VITE_BACKEND_URL}/refine-week-menu?` + params, {
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

export { useRefineWeekMenus };
