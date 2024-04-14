import {
  createProduct,
  getAllProductStock,
  getAllProducts,
  getProduct,
  getProductStock,
  removeProduct,
  upsertProductStock,
} from "../database/products.js";
import { getValueCaseInsensitive } from "../utils/index.js";
import { logger } from "../utils/logger.js";

export const createProductHandler = async (req, res) => {
  const name = getValueCaseInsensitive(req.body, "name");
  const pricing = getValueCaseInsensitive(req.body, "pricing");
  const description = getValueCaseInsensitive(req.body, "description");
  logger.debug(`Product ${name} to be inserted`);
  try {
    const result = await createProduct(name, pricing, description);
    return res.status(200).json({ productId: result });
  } catch (e) {
    logger.error(e);
    return res.status(500).json({ Error: "Error creating product" });
  }
};

export const getAllProductsHandler = async (req, res) => {
  logger.debug("Getting all products");
  try {
    const result = await getAllProducts();
    return res.status(200).json(result);
  } catch (e) {
    logger.error(e);
    return res.status(500).json({ Error: "Error getting products" });
  }
};

export const getProductHandler = async (req, res) => {
  const productId = req.params.id;
  logger.debug(`Getting product with id ${productId}`);
  try {
    const result = await getProduct(productId);
    return res.status(200).json(result);
  } catch (e) {
    logger.error(e);
    return res.status(500).json({ Error: "Error getting product" });
  }
};

export const setProductStockHandler = async (req, res) => {
  const productId = req.params.id;
  const stock = getValueCaseInsensitive(req.body, "quantity");
  logger.debug(`Setting stock for product with id ${productId}`);
  try {
    const result = await upsertProductStock(productId, stock);
    return res.status(200).json(result);
  } catch (e) {
    logger.error(e);
    return res.status(500).json({ Error: "Error setting stock" });
  }
};

export const getAllProductStockHandler = async (req, res) => {
  logger.debug("Getting all product stock");
  try {
    const result = await getAllProductStock();
    return res.status(200).json(result);
  } catch (e) {
    logger.error(e);
    return res.status(500).json({ Error: "Error getting product stock" });
  }
};

export const getProductStockHandler = async (req, res) => {
  const productId = req.params.id;
  logger.debug(`Getting stock for product with id ${productId}`);
  try {
    const result = await getProductStock(productId);
    return res.status(200).json(result[0].stock);
  } catch (e) {
    logger.error(e);
    return res.status(500).json({ Error: "Error getting product stock" });
  }
};

export const removeProductHandler = async (req, res) => {
  const productId = req.params.id;
  logger.debug(`Removing product with id ${productId}`);
  try {
    const result = await removeProduct(productId);
    return res.status(200).json(result);
  } catch (e) {
    logger.error(e);
    return res.status(500).json({ Error: "Error removing product" });
  }
};
