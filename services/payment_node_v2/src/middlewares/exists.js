import { logger } from "../utils/logger.js";

export const productExists = async (req, res, next) => {
  const { productId, quantity } = req.body;

  const result = await fetch(
    process.env.STORE_SERVICE_URL + "/stock/" + productId,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );
  if (!result.ok) {
    logger.error("Error on auth");
    return res.status(401).json({ error: "Unauthorized" });
  }
  const data = await result.json();
  const { quantity: stockQuantity } = data;
  if (stockQuantity < quantity) {
    return res.status(400).json({ error: "Not enough stock" });
  }
  next();
};
