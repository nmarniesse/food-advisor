import { FC } from "react";
import { Box } from "@mui/material";

export const Section: FC<{ children: ReactDOM }> = ({ children }) => {
  return (
    <Box component="section" sx={{ margin: "15px" }}>
      {children}
    </Box>
  );
};
