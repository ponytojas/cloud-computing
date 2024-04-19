import React, { useEffect, useRef, useState } from "react";
import Box from "@mui/material/Box";
import { ProductCard } from "./ProductCard";

export const Products = () => {
  const [products, setProducts] = useState([]);
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
