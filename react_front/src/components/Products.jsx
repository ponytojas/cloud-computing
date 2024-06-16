import React, { useEffect, useRef, useState } from "react";
import Box from "@mui/material/Box";
import { ProductCard } from "./ProductCard";
import { useStore } from "../store";

export const Products = () => {
  const products = useStore((state) => state.products);
  const setProducts = useStore((state) => state.setProducts);
  const fetchedRef = useRef(false);

  useEffect(() => {
    console.debug({ products });
  }, [products]);

  useEffect(() => {
    if (fetchedRef.current) {
      return;
    }
    const getData = async () => {
      const res = await fetch(`${import.meta.env.VITE_CORE_BASE}/v1/stock`);
      const { data } = await res.json();
      // If data is an object with a property called "products" then we set the products to that array
      const _data = data.products ? data.products : data;
      // If _data is an object convert it to an array of the values
      const _products = Array.isArray(_data) ? _data : Object.values(_data);
      console.debug("Type of _products", typeof _products);
      console.debug({ _products });
      setProducts(_products);
    };

    fetchedRef.current = true;
    getData();
  }, []);

  return (
    products &&
    Array.isArray(products) &&
    products.length > 0 && (
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
        {products
          .filter((p) => p.product_id)
          .map((product) => (
            <Box key={product.product_id}>
              <ProductCard product={product} />
            </Box>
          ))}
      </Box>
    )
  );
};
