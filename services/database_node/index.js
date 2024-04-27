import express from "express";
import cors from "cors";
import dotenv from "dotenv";
import { createUserHandler, loginUserHandler } from "./src/handlers/users.js";
import { logger } from "./src/utils/logger.js";
import {
  createProductHandler,
  getAllProductStockHandler,
  getAllProductsHandler,
  getProductHandler,
  getProductStockHandler,
  removeProductHandler,
  setProductStockHandler,
} from "./src/handlers/products.js";

dotenv.config();

const app = express();
app.disable("x-powered-by");
app.use(cors());

app.use(express.json());

app.get("/database/health", (req, res) => {
  res.status(200).send("OK");
});

app.post("/users/create", createUserHandler);
app.post("/users/login", loginUserHandler);

app.post("/products/create", createProductHandler);
app.get("/products", getAllProductsHandler);
app.get("/products/:id", getProductHandler);
app.post("/stock/:id/", setProductStockHandler);
app.get("/stock", getAllProductStockHandler);
app.get("/stock/:id", getProductStockHandler);
app.delete("/products/:id", removeProductHandler);

const port = process.env.HTTP_PORT || 5555;
app.listen(port, () => {
  logger.info(`Server is running on port ${port}`);
});
