import { Router } from "express";
import {} from "../handlers/stock.js";
import { checkToken } from "../middlewares/token.js";
import { getStockHandler } from "../handlers/stock.js";
import { getStockByIdHandler } from "../handlers/stock.js";
import { createStockHandler } from "../handlers/stock.js";

export const getStockRoute = () => {
  const stockRouter = Router();
  stockRouter.get("/", getStockHandler);
  stockRouter.get("/:id", getStockByIdHandler);
  stockRouter.post("/:id", checkToken, createStockHandler);

  return stockRouter;
};
