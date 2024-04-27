import express, { Router } from "express";
import cors from "cors";
import dotenv from "dotenv";
import listEndpoints from "express-list-endpoints";

import { logger } from "./src/utils/logger.js";
import { getUsersRoute } from "./src/routes/users.js";
import { getProductsRoute } from "./src/routes/products.js";
import { getStockRoute } from "./src/routes/stock.js";

dotenv.config();

const app = express();
app.disable("x-powered-by");
app.use(cors());
app.use(express.json());

app.use((req, res, next) => {
  logger.debug(
    `[${new Date().toISOString()}] ${req.method} ${req.originalUrl}`
  );
  next();
});

const apiV1Router = Router();
const userRouter = getUsersRoute();
const productRouter = getProductsRoute();
const stockRouter = getStockRoute();

apiV1Router.get("/health", (req, res) => {
  res.status(200).send("OK");
});

apiV1Router.use("/user", userRouter);
apiV1Router.use("/product", productRouter);
apiV1Router.use("/stock", stockRouter);

app.use("/v1", apiV1Router);

const port = process.env.HTTP_PORT || 5550;
const server = app.listen(port, () => {
  logger.info(`Server is running on port ${port}`);
  const endpoints = listEndpoints(app);
  endpoints.forEach((endpoint) => {
    logger.info(`${endpoint.methods.join(", ")} ${endpoint.path}`);
  });
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
