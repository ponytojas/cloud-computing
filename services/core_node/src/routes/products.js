import { Router } from "express";
import {
  createProductHandler,
  deleteProductHandler,
  getProductByIdHandler,
  getProductsHandler,
} from "../handlers/products.js";
import { checkToken } from "../middlewares/token.js";

export const getProductsRoute = () => {
  const productRouter = Router();
  productRouter.get("/", getProductsHandler);
  productRouter.get("/:id", getProductByIdHandler);
  productRouter.post("/", checkToken, createProductHandler);
  productRouter.delete("/:id", checkToken, deleteProductHandler);

  return productRouter;
};
