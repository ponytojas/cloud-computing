import { logger } from "../utils/logger.js";
import { Redis } from "./redis.js";
import * as jose from "jose";
const secret = new TextEncoder().encode(process.env.JWT_SECRET);

export const createToken = async ({ userId, username, email }) => {
  const redisClient = new Redis();
  let token = await redisClient.get(userId);
  if (token) {
    logger.debug(`Token for user ${userId} found in cache: ${token}`);
    return token;
  }
  const payload = {
    userId,
    username,
    email,
  };
  logger.debug(`Creating token for user ${userId}`);
  token = await new jose.SignJWT(payload)
    .setProtectedHeader({ alg: "HS256" })
    .setIssuedAt()
    .setExpirationTime("24h")
    .sign(secret);
  await redisClient.set(userId, token);
  return token;
};

export const checkToken = async (token) => {
  try {
    const { payload } = await jose.jwtVerify(token, secret);
    if (payload.exp < Date.now() / 1000) {
      logger.debug("Token expired");
      return null;
    }
    const redisClient = new Redis();
    const cachedToken = await redisClient.get(payload.userId);
    if (cachedToken !== token) {
      logger.debug("Token not found in cache");
      return null;
    }
    return { claims: payload, valid: true };
  } catch (error) {
    logger.error(`Error verifying token: ${error}`);
    return null;
  }
};

export const deleteToken = async (userId) => {
  const redisClient = new Redis();
  await redisClient.del(userId);
};
