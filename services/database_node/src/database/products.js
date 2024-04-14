import { Database } from "./db.js";

export const createProduct = async (name, pricing, description) => {
  const client = new Database();
  const createdAt = new Date();
  const query = `INSERT INTO "product" (name, pricing, description, created_at) VALUES ($1, $2, $3, $4) RETURNING product_id`;
  const values = [name, pricing, description, createdAt];
  const result = await client.query(query, values);
  return result.rows[0].product_id;
};

export const getAllProducts = async () => {
  const client = new Database();
  const query = `SELECT * FROM "product"`;
  const result = await client.query(query);
  return result.rows;
};

export const getProduct = async (productId) => {
  const client = new Database();
  const query = `SELECT * FROM "product" WHERE product_id = $1`;
  const values = [productId];
  const result = await client.query(query, values);
  return result.rows;
};

export const upsertProductStock = async (productId, stock) => {
  const client = new Database();
  const query = `INSERT INTO "product_stock" (product_id, quantity) VALUES ($1, $2) ON CONFLICT (product_id) DO UPDATE SET quantity = EXCLUDED.quantity;`;
  const values = [productId, stock];
  await client.query(query, values);
  return { productId, stock };
};

export const getAllProductStock = async () => {
  const client = new Database();
  const query = `SELECT * FROM "product" p JOIN "product_stock" ps ON p.product_id = ps.product_id`;
  const result = await client.query(query);
  return result.rows;
};

export const getProductStock = async (productId) => {
  const client = new Database();
  const query = `SELECT * FROM "product" p JOIN "product_stock" ps ON p.product_id = ps.product_id WHERE p.product_id = $1`;
  const values = [productId];
  const result = await client.query(query, values);
  return result.rows;
};

export const removeProduct = async (productId) => {
  const client = new Database();
  const query = `DELETE FROM "product" WHERE product_id = $1`;
  const values = [productId];
  await client.query(query, values);
  const query2 = `DELETE FROM "product_stock" WHERE product_id = $1`;
  await client.query(query2, values);
  return { productId };
};
