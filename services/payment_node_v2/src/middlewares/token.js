import { logger } from "../utils/logger.js";

export const checkToken = async (req, res, next) => {
  const { authorization } = req?.headers || null;
  if (!authorization) return res.status(401).json({ error: "Unauthorized" });
  const body = JSON.stringify({ token: authorization });

  const result = await fetch(process.env.AUTH_SERVICE_URL + "/check", {
    method: "POST",
    body,
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (result.status !== 200) {
    logger.error("Error on auth");
    return res.status(401).json({ error: "Unauthorized" });
  }
  next();
};
