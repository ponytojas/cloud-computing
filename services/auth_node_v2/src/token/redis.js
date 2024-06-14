import { createClient } from "redis";
import { logger } from "../utils/logger.js";

export class Redis {
  static instance;

  constructor() {
    if (!Redis.instance) {
      this.client = createClient({
        url: `redis://${process.env.REDIS_HOST}:${process.env.REDIS_PORT}`,
      });

      Redis.instance = this;
      this.connectWithRetry();
    }

    return Redis.instance;
  }

  // Method to handle connection with retry mechanism
  connectWithRetry(maxAttempts = 25, interval = 2000) {
    let attempts = 0;

    const connect = async () => {
      try {
        await this.client.connect();
        logger.info("Successfully connected to Redis");
      } catch (error) {
        attempts++;
        if (error.message.includes("Socket already opened")) {
          logger.error(
            `Connection attempt ${attempts} failed: ${error.message}`
          );
          // Disconnect the client if the socket is already open
          await this.client.disconnect();
        } else {
          logger.error(
            `Connection attempt ${attempts} failed: ${error.message}`
          );
        }

        if (attempts < maxAttempts) {
          setTimeout(connect, interval);
        } else {
          logger.error(
            "Failed to connect to Redis after several attempts",
            error
          );
        }
      }
    };

    connect();
  }

  // Disconnect from Redis
  static disconnect() {
    if (Redis.instance) {
      Redis.instance.client.disconnect();
      Redis.instance = null;
    }
  }

  async _set(key, value) {
    const _key = typeof key === "string" ? key : JSON.stringify(key);
    try {
      await this.client.set(_key, value, { EX: 60 * 60 * 24 });
    } catch (e) {
      logger.error(`Error setting key ${_key}: ${e.message}`);
      return null;
    }
  }

  async _get(key) {
    console.log("Getting key:", key);
    const _key = typeof key === "string" ? key : JSON.stringify(key);
    try {
      const reply = await this.client.get(_key);
      return reply;
    } catch (e) {
      logger.error(`Error getting key ${_key}: ${e.message}`);
      return null;
    }
  }

  async _del(key) {
    const _key = typeof key === "string" ? key : JSON.stringify(key);
    try {
      await this.client.del(_key);
    } catch (e) {
      logger.error(`Error deleting key ${_key}: ${e.message}`);
      return null;
    }
  }
}

// Ensure that the singleton instance is available immediately
export const redisClient = new Redis().client;
