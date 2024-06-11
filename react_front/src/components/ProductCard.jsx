import React from "react";
import PropTypes from "prop-types";

import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Rating from "@mui/material/Rating";
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
        padding: 4,
      }}
    >
      {product.picture && (
        <Box
          sx={{
            display: "flex",
            marginTop: 2,
            marginBottom: 2,
            width: "100%",
            justifyContent: "center",
            alignContent: "center",
            alignItems: "center",
          }}
        >
          <img
            src="https://ugmonk.com/cdn/shop/files/new-magsafe-white-maple-2_1480x1836_crop_center.jpg?v=1714765834"
            alt={product.name}
            style={{ width: "128px", margin: "0 auto" }}
          />
        </Box>
      )}
      <Typography variant="h6">{product.name}</Typography>
      <Typography>{product.description}</Typography>
      <Typography>${product.pricing}</Typography>
      <Typography>Stock: {product.quantity}</Typography>
      {product.rating && (
        <Rating name="read-only" value={product.rating} readOnly />
      )}
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
