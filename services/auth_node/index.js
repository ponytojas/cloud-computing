import express from "express";
import cors from "cors";
import dotenv from "dotenv";
import { escape } from "html-escaper";

import {
  registerHandler,
  loginUserHandler,
  logoutUserHandler,
} from "./src/handlers/users.js";
import { logger } from "./src/utils/logger.js";
import { checkToken } from "./src/token/token.js";
import { redisClient } from "./src/token/redis.js";

dotenv.config();

const app = express();
app.disable("x-powered-by");

app.use(cors());

app.use(express.json());

app.use((req, res, next) => {
  console.log("Headers:", JSON.stringify(req.headers));
  console.log("Body:", JSON.stringify(req.body));
  next();
});

app.get("/auth/health", (req, res) => {
  res.status(200).send("OK");
});

app.post("/register", registerHandler);
app.post("/login", loginUserHandler);
app.post("/logout", logoutUserHandler);
app.post("/check", async (req, res) => {
  const { token } = req.body;
  console.log("Token:", token);
  if (token) {
    const result = await checkToken(token);
    if (result) {
      res.status(200).send(escape(result));
    } else {
      res.status(401).send("Unauthorized");
    }
  } else {
    res.status(401).send("Unauthorized");
  }
});

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
