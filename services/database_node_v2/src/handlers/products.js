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
  const now = new Date().getTime();
  const name = getValueCaseInsensitive(req.body, "name");
  const pricing = getValueCaseInsensitive(req.body, "pricing");
  const description = getValueCaseInsensitive(req.body, "description");
  const rating = getValueCaseInsensitive(req.body, "rating");
  const picture = getValueCaseInsensitive(req.body, "picture");
  logger.debug(`Product ${name} to be inserted`);
  try {
    const result = await createProduct(
      name,
      pricing,
      description,
      rating,
      picture
    );
    const end = new Date().getTime();
    return res.status(200).json({ productId: result, time: end - now });
  } catch (e) {
    logger.error(e);
    return res.status(500).json({ Error: "Error creating product" });
  }
};

export const getAllProductsHandler = async (req, res) => {
  logger.debug("Getting all products");
  try {
    const now = new Date().getTime();
    const result = await getAllProducts();
    const end = new Date().getTime();
    return res.status(200).json({ ...result, time: end - now });
  } catch (e) {
    logger.error(e);
    return res.status(500).json({ Error: "Error getting products" });
  }
};

export const getProductHandler = async (req, res) => {
  const now = new Date().getTime();
  const productId = req.params.id;
  logger.debug(`Getting product with id ${productId}`);
  try {
    const result = await getProduct(productId);
    const end = new Date().getTime();
    return res.status(200).json({ ...result, time: end - now });
  } catch (e) {
    logger.error(e);
    return res.status(500).json({ Error: "Error getting product" });
  }
};

export const setProductStockHandler = async (req, res) => {
  const now = new Date().getTime();
  const productId = req.params.id;
  const stock = getValueCaseInsensitive(req.body, "quantity");
  logger.debug(`Setting stock for product with id ${productId}`);
  try {
    const result = await upsertProductStock(productId, stock);
    const end = new Date().getTime();
    return res.status(200).json({ ...result, time: end - now });
  } catch (e) {
    logger.error(e);
    return res.status(500).json({ Error: "Error setting stock" });
  }
};

export const getAllProductStockHandler = async (req, res) => {
  const now = new Date().getTime();
  logger.debug("Getting all product stock");
  try {
    const result = await getAllProductStock();
    const end = new Date().getTime();
    return res.status(200).json({ data: result, time: end - now });
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
    const end = new Date().getTime();
    return res.status(200).json(result[0]);
  } catch (e) {
    logger.error(e);
    return res.status(500).json({ Error: "Error getting product stock" });
  }
};

export const removeProductHandler = async (req, res) => {
  const now = new Date().getTime();
  const productId = req.params.id;
  logger.debug(`Removing product with id ${productId}`);
  try {
    const result = await removeProduct(productId);
    const end = new Date().getTime();
    return res.status(200).json({ ...result, time: end - now });
  } catch (e) {
    logger.error(e);
    return res.status(500).json({ Error: "Error removing product" });
  }
};
