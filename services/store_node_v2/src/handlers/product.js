import { logger } from "../utils/logger.js";

export const createHandle = async (req, res) => {
  try {
    const now = new Date().getTime();
    const product = req.body;
    const result = await fetch(
      process.env.DB_SERVICE_URL + "/products/create",
      {
        method: "POST",
        body: JSON.stringify(product),
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    const body = await result.json();
    logger.debug(`Product registered with ID: ${body.productId}`);
    const end = new Date().getTime();
    res.status(201).json({ ...body, time: end - now });
  } catch (error) {
    logger.error("Error registering product:", error);
    res.status(500).send("ERROR 1001");
  }
};

export const getHandle = async (req, res) => {
  try {
    const now = new Date().getTime();
    const result = await fetch(process.env.DB_SERVICE_URL + "/products", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const body = await result.json();
    const end = new Date().getTime();
    res.status(200).json({ ...body, time: end - now });
  } catch (error) {
    logger.error("Error fetching products:", error);
    res.status(500).send("ERROR 1002");
  }
};

export const getByIdHandle = async (req, res) => {
  try {
    const now = new Date().getTime();
    const id = req.params.id;
    const result = await fetch(process.env.DB_SERVICE_URL + `/products/${id}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const body = await result.json();
    const end = new Date().getTime();
    res.status(200).json({ ...body, time: end - now });
  } catch (error) {
    logger.error("Error fetching product by ID:", error);
    res.status(500).send("ERROR 1003");
  }
};

export const addStockHandle = async (req, res) => {
  try {
    const now = new Date().getTime();
    const id = req.params.id;
    const stock = req.body;
    const result = await fetch(process.env.DB_SERVICE_URL + `/stock/${id}`, {
      method: "POST",
      body: JSON.stringify(stock),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const body = await result.json();
    const end = new Date().getTime();
    res.status(200).json({ ...body, time: end - now });
  } catch (error) {
    logger.error("Error adding stock to product:", error);
    res.status(500).send("ERROR 1004");
  }
};

export const getStockHandle = async (req, res) => {
  try {
    const now = new Date().getTime();
    const id = req.params.id;
    const result = await fetch(process.env.DB_SERVICE_URL + `/stock/${id}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const body = await result.json();
    const end = new Date().getTime();
    res.status(200).json({ ...body, time: end - now });
  } catch (error) {
    logger.error("Error fetching stock by product ID:", error);
    res.status(500).send("ERROR 1005");
  }
};

export const getAllStockHandle = async (req, res) => {
  try {
    const now = new Date().getTime();
    const result = await fetch(process.env.DB_SERVICE_URL + "/stock", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const body = await result.json();
    const end = new Date().getTime();
    res.status(200).json({ ...body, time: end - now });
  } catch (error) {
    logger.error("Error fetching all stock:", error);
    res.status(500).send("ERROR 1006");
  }
};

export const deleteHandle = async (req, res) => {
  try {
    const now = new Date().getTime();
    const id = req.params.id;
    const result = await fetch(process.env.DB_SERVICE_URL + `/products/${id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const body = await result.json();
    const end = new Date().getTime();
    res.status(200).json({ ...body, time: end - now });
  } catch (error) {
    logger.error("Error deleting product:", error);
    res.status(500).send("ERROR 1007");
  }
};
