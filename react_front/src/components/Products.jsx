import React, { useEffect, useRef, useState } from "react";
import Box from "@mui/material/Box";
import { ProductCard } from "./ProductCard";
import { useStore } from "../store";

export const Products = () => {
  const products = useStore((state) => state.products);
  const setProducts = useStore((state) => state.setProducts);
  const fetchedRef = useRef(false);

  useEffect(() => {
    if (fetchedRef.current) {
      return;
    }
    const getData = async () => {
      const res = await fetch(`${import.meta.env.VITE_CORE_BASE}/v1/stock`);
      const data = await res.json();
      setProducts(data);
    };

    fetchedRef.current = true;
    getData();
  }, []);
  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "row",
        flexWrap: "wrap",
        alignContent: "center",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      {products.map((product) => (
        <Box key={product.product_id}>
          <ProductCard product={product} />
        </Box>
      ))}
    </Box>
  );
};
