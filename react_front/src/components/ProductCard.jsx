import React from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

export const ProductCard = ({ product }) => {
  return (
    <Box
      sx={{
        p: 2,
        m: 1,
        border: "1px solid #ccc",
        minWidth: "250px",
        borderRadius: 1,
      }}
    >
      <Typography variant="h6">{product.name}</Typography>
      <Typography>{product.description}</Typography>
      <Typography>${product.pricing}</Typography>
    </Box>
  );
};
