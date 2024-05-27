import React, { useEffect, useState } from "react";
import PropTypes from "prop-types";
import { Box, Button, Drawer, Typography } from "@mui/material";
import CloseFullscreenIcon from "@mui/icons-material/CloseFullscreen";
import { useStore } from "../store";

export const BottomDrawer = ({ open, toggleDrawer }) => {
  const [totalPrice, setTotalPrice] = useState(0);

  const cart = useStore((state) => state.cart);
  const setCart = useStore((state) => state.setCart);

  const token = useStore((state) => state.token);
  const userId = useStore((state) => state.userId);

  useEffect(() => {
    let total = 0;
    Object.values(cart).forEach((v) => (total += v.price * v.quantity));
    setTotalPrice(total);
  }, [cart]);

  const handleBuy = async () => {
    try {
      await fetch(`${import.meta.env.VITE_PAYMENT_BASE}/v1/pay/${userId}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: token,
        },
      });
      setCart({});
    } catch (e) {
      console.error(e);
    }
  };

  return (
    open && (
      <Box sx={{ height: "300px" }}>
        <Drawer
          open={open}
          onClose={() => {
            toggleDrawer();
          }}
          anchor={"bottom"}
          sx={{ minHeight: "300px" }}
        >
          <Box
            sx={{
              display: "flex",
              width: "100%",
              justifyContent: "flex-end",
            }}
          >
            <CloseFullscreenIcon
              onClick={toggleDrawer}
              sx={{ cursor: "pointer", mr: 3, mt: 2 }}
            />
          </Box>
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              height: "300px",
              overflowX: "auto",
              p: 4,
              justifyContent: "center",
            }}
          >
            <Box
              sx={{
                display: "flex",
                flexDirection: "row",
                justifyContent: "center",
              }}
            >
              <Typography variant="h3">{`Total price: ${totalPrice}€`}</Typography>
            </Box>
            {totalPrice > 0 && (
              <Box
                sx={{
                  display: "flex",
                  flexDirection: "row",
                  width: "auto",
                  justifyContent: "center",
                  mt: 4,
                }}
              >
                <Button
                  variant="contained"
                  sx={{ px: 6, py: 1, borderRadius: "20px" }}
                  onClick={handleBuy}
                >
                  Buy
                </Button>
              </Box>
            )}
          </Box>
        </Drawer>
      </Box>
    )
  );
};

BottomDrawer.propTypes = {
  open: PropTypes.bool,
  toggleDrawer: PropTypes.func,
};
