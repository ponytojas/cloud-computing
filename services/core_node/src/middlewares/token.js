import axios from "axios";
import { logger } from "../utils/logger.js";

export const checkToken = async (req, res, next) => {
  try {
    const { authorization } = req?.headers || null;
    if (!authorization)
      return res.status(401).json({ error: "Unauthorized - Not header found" });

    const result = await axios.post(process.env.AUTH_SERVICE_URL + "/check", {
      token: authorization,
    });

    if (result.status !== 200) {
      logger.error("Error on auth");
      return res.status(401).json({ error: "Unauthorized" });
    }
  } catch (error) {
    logger.error(`Error verifying token: ${error}`);
    return res.status(401).json({ error: "Unauthorized" });
  }
  next();
};
