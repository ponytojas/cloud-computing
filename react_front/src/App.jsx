import { Box, IconButton, Typography } from "@mui/material";
import MenuIcon from "@mui/icons-material/Menu";
import React, { useEffect, useRef, useState } from "react";
import { LoginDrawer } from "./components/Drawer";
import { Products } from "./components/Products";

export const App = () => {
  const [open, setOpen] = useState(false);
  const toggleDrawer = () => () => {
    setOpen((prev) => !prev);
  };

  return (
    <>
      <LoginDrawer open={open} toggleDrawer={toggleDrawer} />
      <Box
        sx={{
          display: "flex",
          height: "100vh",
          width: "100vw",
          flexDirection: "column",
          alignItems: "center",
          p: 2,
        }}
      >
        <Box
          sx={{
            display: "flex",
            flexDirection: "row",
            width: "100%",
            justifyContent: "space-between",
          }}
        >
          <Typography variant="h2" sx={{ fontWeight: 100 }}>
            Products
          </Typography>
          <IconButton aria-label="delete" onClick={toggleDrawer(true)}>
            <MenuIcon />
          </IconButton>
        </Box>

        <Box>
          <Products />
        </Box>
      </Box>
    </>
  );
};
