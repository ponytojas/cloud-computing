import { Box, Button, IconButton, Typography } from "@mui/material";
import MenuIcon from "@mui/icons-material/Menu";
import ShoppingCartIcon from "@mui/icons-material/ShoppingCart";
import React, { useCallback, useEffect, useState } from "react";
import { LoginDrawer } from "./components/Drawer";
import { Products } from "./components/Products";
import { useStore } from "./store";
import { BottomDrawer } from "./components/BottomDrawer";

export const App = () => {
  const [open, setOpen] = useState(false);
  const [openCart, setOpenCart] = useState(false);
  const [items, setItems] = useState(0);

  const cart = useStore((state) => state.cart);
  const token = useStore((state) => state.token);

  const toggleDrawer = () => {
    setOpen((prev) => !prev);
  };

  const toggleBottomDrawer = useCallback(() => {
    setOpenCart((prev) => !prev);
  }, [openCart]);

  useEffect(() => {
    let total = 0;
    Object.values(cart).forEach((v) => (total += v.quantity));
    setItems(total);
  }, [cart]);

  return (
    <>
      {token && (
        <Box
          sx={{
            position: "fixed",
            bottom: "3vh",
            width: "100%",
            p: 2,
            zIndex: 1000,
          }}
        >
          <Box
            sx={{ display: "flex", width: "100%", justifyContent: "center" }}
          >
            <Button
              color="success"
              sx={{
                borderRadius: "20px",
                px: 4,
                py: 1,
                boxShadow:
                  " 0 3px 6px rgba(0,0,0,0.16), 0 3px 6px rgba(0,0,0,0.23)",
              }}
              variant="contained"
              startIcon={<ShoppingCartIcon sx={{ mr: 1 }} />}
              onClick={toggleBottomDrawer}
            >
              {`Cart (${items})`}
            </Button>
          </Box>
        </Box>
      )}
      <BottomDrawer open={openCart} toggleDrawer={toggleBottomDrawer} />

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
          <IconButton aria-label="delete" onClick={toggleDrawer}>
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
