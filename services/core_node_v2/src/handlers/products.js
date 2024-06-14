import { adaptResponse } from "../utils/body.js";
import { logger } from "../utils/logger.js";

export const createProductHandler = async (req, res) => {
  if (!req.body) return res.status(500).json({ error: "Missing request body" });

  try {
    const result = await fetch(process.env.STORE_SERVICE_URL + "/product", {
      method: "POST",
      body: JSON.stringify(req.body),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const body = await result.json();
    if (!result.ok) {
      logger.error("Error creating product:", body);
      return res.status(500).send("ERROR 1002");
    }
    return res.status(200).json({ ...body });
  } catch (error) {
    logger.error("Error creating product:", error);
    return res.status(500).send("ERROR 1002");
  }
};

export const getProductsHandler = async (req, res) => {
  try {
    const result = await fetch(process.env.STORE_SERVICE_URL + "/product");
    const body = await result.json();
    const _body = adaptResponse(body);
    if (!result.ok) {
      logger.error("Error getting products:", body);
      return res.status(500).send("ERROR 1003");
    }
    return res.status(200).json(_body);
  } catch (error) {
    logger.error("Error getting products:", error);
    return res.status(500).send("ERROR 1003");
  }
};

export const getProductByIdHandler = async (req, res) => {
  try {
    const result = await fetch(
      process.env.STORE_SERVICE_URL + "/product/" + req.params.id
    );
    const body = await result.json();
    if (!result.ok) {
      logger.error("Error getting product:", body);
      return res.status(500).send("ERROR 1004");
    }
    return res.status(200).json(body);
  } catch (error) {
    logger.error("Error getting product:", error);
    return res.status(500).send("ERROR 1004");
  }
};

export const deleteProductHandler = async (req, res) => {
  try {
    const result = await fetch(
      process.env.STORE_SERVICE_URL + "/products/" + req.params.id,
      {
        method: "DELETE",
      }
    );
    const body = await result.json();
    if (!result.ok) {
      logger.error("Error deleting product:", body);
      return res.status(500).send("ERROR 1005");
    }
    return res.status(200).json(body);
  } catch (error) {
    logger.error("Error deleting product:", error);
    return res.status(500).send("ERROR 1005");
  }
};
