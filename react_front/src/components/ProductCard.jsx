import React from "react";
import PropTypes from "prop-types";

import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { Button } from "@mui/material";
import { toast } from "react-toastify";
import { useStore } from "../store";

export const ProductCard = ({ product }) => {
  const cart = useStore((state) => state.cart);
  const setCart = useStore((state) => state.setCart);
  const userId = useStore((state) => state.userId);
  const token = useStore((state) => state.token);

  const handleClick = async () => {
    if (product.quantity > 0) {
      const newEntry = cart[product.product_id]
        ? {
            ...cart[product.product_id],
            quantity: cart[product.product_id].quantity + 1,
          }
        : {
            productId: product.product_id,
            quantity: 1,
            price: product.pricing,
          };

      setCart({
        ...cart,
        [product.product_id]: newEntry,
      });

      const data = {
        userId,
        productId: product.product_id,
        quantity: newEntry.quantity,
      };

      try {
        await fetch(`${import.meta.env.VITE_CART_BASE}/v1/add-to-cart`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
          body: JSON.stringify(data),
        });
      } catch (e) {
        console.error(e);
      }
    } else {
      toast.error("Not enough stock");
    }
  };

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
      <Typography>Stock: {product.quantity}</Typography>
      {token && (
        <Button
          variant="contained"
          color="primary"
          sx={{ mt: 2 }}
          onClick={handleClick}
        >
          Add to Cart
        </Button>
      )}
    </Box>
  );
};

ProductCard.propTypes = {
  product: PropTypes.shape({
    product_id: PropTypes.number,
    name: PropTypes.string,
    description: PropTypes.string,
    pricing: PropTypes.string,
    quantity: PropTypes.number,
  }),
};
