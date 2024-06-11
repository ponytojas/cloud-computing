import { logger } from "../utils/logger.js";

export const getStockHandler = async (req, res) => {
  try {
    const result = await fetch(process.env.STORE_SERVICE_URL + "/stock");
    const body = await result.json();
    if (!result.ok) {
      logger.error("Error getting products:", body);
      return res.status(500).send("ERROR 1003");
    }
    return res.status(200).json(body);
  } catch (error) {
    logger.error("Error getting products:", error);
    return res.status(500).send("ERROR 1003");
  }
};

export const getStockByIdHandler = async (req, res) => {
  try {
    const result = await fetch(
      `${process.env.STORE_SERVICE_URL}/stock/${req.params.id}`
    );
    const body = await result.json();
    if (!result.ok) {
      logger.error("Error getting product:", body);
      return res.status(500).send("ERROR 1003");
    }
    return res.status(200).json(body);
  } catch (error) {
    logger.error("Error getting product:", error);
    return res.status(500).send("ERROR 1003");
  }
};

export const createStockHandler = async (req, res) => {
  try {
    if (!req.params.id) return res.status(400).send("ERROR 1001");
    const result = await fetch(
      `${process.env.STORE_SERVICE_URL}/stock/${req.params.id}`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: req.headers.authorization,
        },
        body: JSON.stringify(req.body),
      }
    );
    const body = await result.json();
    if (!result.ok) {
      logger.error("Error creating product:", body);
      return res.status(500).send("ERROR 1003");
    }
    return res.status(201).json(body);
  } catch (error) {
    logger.error("Error creating product:", error);
    return res.status(500).send("ERROR 1003");
  }
};
