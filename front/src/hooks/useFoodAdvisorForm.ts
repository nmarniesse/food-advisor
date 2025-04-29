import { MenuData } from "../models/MenuData";

type Return = {
  submit: (menuData: MenuData) => void;
};

const useFoodAdvisorForm = (): Return => {
  const submit = (menuData: MenuData) => {
    console.log(menuData);
    console.log("VITE_BACKEND_URL", import.meta.env.VITE_BACKEND_URL);
  };

  return { submit };
};

export { useFoodAdvisorForm };
