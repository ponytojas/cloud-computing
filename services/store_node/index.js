import express from "express";
import cors from "cors";
import dotenv from "dotenv";
import {
  addStockHandle,
  createHandle,
  deleteHandle,
  getAllStockHandle,
  getByIdHandle,
  getHandle,
  getStockHandle,
} from "./src/handlers/product.js";
import { logger } from "./src/utils/logger.js";

dotenv.config();

const app = express();
app.disable("x-powered-by");
app.use(cors());

app.use(express.json());

app.get("/auth/health", (req, res) => {
  res.status(200).send("OK");
});

app.post("/product", createHandle);
app.get("/product", getHandle);
app.get("/product/:id", getByIdHandle);
app.post("/stock/:id", addStockHandle);
app.get("/stock/:id", getStockHandle);
app.get("/stock", getAllStockHandle);
app.delete("/products/:id", deleteHandle);

const port = process.env.HTTP_PORT || 5554;
const server = app.listen(port, () => {
  logger.info(`Server is running on port ${port}`);
});

process.on("SIGINT", gracefulShutdown);
process.on("SIGTERM", gracefulShutdown);

function gracefulShutdown() {
  logger.info("Received shutdown signal. Initiating graceful shutdown...");

  // Stop accepting new connections
  server.close(() => {
    logger.info("Server closed. No longer accepting new connections.");
    logger.info("Graceful shutdown complete. Exiting.");
    process.exit(0);
  });

  setTimeout(() => {
    logger.warn("Graceful shutdown timeout exceeded. Forcing exit.");
    process.exit(1);
  }, 5000); // Adjust the timeout as needed
}
