import express, { Router } from "express";
import cors from "cors";
import dotenv from "dotenv";
import listEndpoints from "express-list-endpoints";

import { logger } from "./src/utils/logger.js";
import { checkToken } from "./src/middlewares/token.js";
import { handlePay, handleTotal } from "./src/handlers/pay.js";
import { redisClient } from "./src/utils/redis.js";

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

apiV1Router.get("/health", (req, res) => {
  res.status(200).send("OK");
});

apiV1Router.get("/pay/:userId", checkToken, handlePay);
apiV1Router.get("/total/:userId", checkToken, handleTotal);

app.use("/v1", apiV1Router);

const port = process.env.HTTP_PORT || 5560;
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
    redisClient.disconnect();
    logger.info("Server closed. No longer accepting new connections.");
    logger.info("Graceful shutdown complete. Exiting.");
    process.exit(0);
  });

  setTimeout(() => {
    logger.warn("Graceful shutdown timeout exceeded. Forcing exit.");
    process.exit(1);
  }, 5000); // Adjust the timeout as needed
}
