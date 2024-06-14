import { Redis } from "../utils/redis.js";

export const handleAdd = async (req, res) => {
  const now = new Date().getTime();
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
  const end = new Date().getTime();
  res.status(200).send({ status: "OK", time: end - now });
};

export const handleGet = async (req, res) => {
  const now = new Date().getTime();
  const redisClient = new Redis();
  const { userId } = req.params;
  const key = `cart-user-${userId}`;
  const cart = await redisClient._get(key);
  if (!cart) {
    return res.status(404).json({ error: "Cart not found" });
  }
  const end = new Date().getTime();
  const result = { status: "OK", time: end - now, cart: JSON.parse(cart) };
  res.status(200).send(result);
};

export const handleDelete = async (req, res) => {
  const now = new Date().getTime();
  const redisClient = new Redis();
  const { userId } = req.params;
  const key = `cart-user-${userId}`;
  const cart = await redisClient._get(key);
  if (!cart) {
    return res.status(404).json({ error: "Cart not found" });
  }
  await redisClient._del(key);
  const end = new Date().getTime();
  res.status(200).send({ status: "OK", time: end - now });
};
