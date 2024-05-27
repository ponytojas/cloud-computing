import { Redis } from "../utils/redis.js";

export const handlePay = async (req, res) => {
  const client = new Redis();
  const { userId } = req.params;

  // Delete the cart from the key `cart-user-userId`
  const cartKey = `cart-user-${userId}`;
  await client._del(cartKey);

  res.status(200).json({ message: "Payment successful" });
};

export const handleTotal = async (req, res) => {
  const client = new Redis();
  const { userId } = req.params;

  // Get the cart from the key `cart-user-userId`
  const cartKey = `cart-user-${userId}`;
  const cart = JSON.parse(await client._get(cartKey));

  if (!cart) {
    return res.status(200).json({});
  }
  const total = await Promise.all(
    Object.entries(cart).map(async ([productId, quantity]) => {
      const response = await fetch(
        `${process.env.STORE_SERVICE_URL}/product/${productId}`
      );
      const product = await response.json();
      const { pricing } = product[0];
      return pricing * quantity;
    })
  );

  const grandTotal = total.reduce((acc, curr) => acc + curr, 0);
  res.status(200).json({ total: grandTotal });
};
