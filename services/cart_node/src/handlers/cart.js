import { Redis } from "../utils/redis.js";

export const handleAdd = async (req, res) => {
  const redisClient = new Redis();
  const { userId, productId, quantity } = req.body;
  // Check if the userId has a cart in redis and the product is already in the cart
  const key = `cart-user-${userId}`;
  const cart = await redisClient._get(key);
  if (!cart) {
    await redisClient._set(key, JSON.stringify({ [productId]: quantity }));
  } else {
    const cartObj = JSON.parse(cart);
    if (cartObj[productId]) {
      cartObj[productId] += quantity;
    } else {
      cartObj[productId] = quantity;
    }
    await redisClient._set(key, JSON.stringify(cartObj));
  }
  res.status(200).send("OK");
};

export const handleGet = async (req, res) => {
  const redisClient = new Redis();
  const { userId } = req.params;
  const key = `cart-user-${userId}`;
  const cart = await redisClient._get(key);
  if (!cart) {
    return res.status(404).json({ error: "Cart not found" });
  }
  res.status(200).json(JSON.parse(cart));
};

export const handleDelete = async (req, res) => {
  const redisClient = new Redis();
  const { userId } = req.params;
  const key = `cart-user-${userId}`;
  const cart = await redisClient._get(key);
  if (!cart) {
    return res.status(404).json({ error: "Cart not found" });
  }
  await redisClient._del(key);
  res.status(200).send("OK");
};
