import axios from "axios";
import { logger } from "../utils/logger.js";

export const checkToken = async (req, res, next) => {
  try {
    const headers = req?.headers || null;
    if (!headers) {
      return res.status(401).json({ error: "Unauthorized - Not header found" });
    }
    const auth = headers.authorization || headers.Authorization || null;
    if (!auth)
      return res.status(401).json({ error: "Unauthorized - Not header found" });

    const result = await axios.post(process.env.AUTH_SERVICE_URL + "/check", {
      token: auth,
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
