import { Router } from "express";
import { checkToken } from "../middlewares/token.js";
import {
  getStockHandler,
  getStockByIdHandler,
  createStockHandler,
} from "../handlers/stock.js";

export const getStockRoute = () => {
  const stockRouter = Router();
  stockRouter.get("/", getStockHandler);
  stockRouter.get("/:id", getStockByIdHandler);
  stockRouter.post("/:id", checkToken, createStockHandler);

  return stockRouter;
};
