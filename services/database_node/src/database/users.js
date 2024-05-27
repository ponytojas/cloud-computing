import bcrypt from "bcryptjs";
import { Database } from "./db.js";
import { logger } from "../utils/logger.js";

export const createUser = async (user) => {
  const hashedPassword = await bcrypt.hash(
    user.password,
    bcrypt.genSaltSync(10)
  );

  const client = new Database();
  const query = `
    INSERT INTO users (username, email, created_at)
    VALUES ($1, $2, $3)
    RETURNING user_id
  `;
  const values = [user.username, user.email, new Date()];
  const result = await client.query(query, values);
  const userId = result.rows[0].user_id;
  await client.query(
    `INSERT INTO auth (user_id, password_hash) VALUES ($1, $2)`,
    [userId, hashedPassword]
  );
  return userId;
};

export const loginUser = async (_username, password) => {
  const client = new Database();
  const query = `
    SELECT auth.user_id, auth.password_hash, users.username, users.email
    FROM users
    INNER JOIN auth ON auth.user_id = users.user_id
    WHERE username=$1 OR email=$1
  `;
  const values = [_username];
  const result = await client.query(query, values);
  logger.info(`Getted user: ${JSON.stringify(result.rows[0])}`);
  const { user_id, password_hash, username, email } = result.rows[0];
  const isPasswordCorrect = await bcrypt.compare(password, password_hash);
  if (!isPasswordCorrect) {
    throw new Error("Incorrect password");
  }
  const returnObj = { userId: user_id, username, email };
  logger.info(`Logged in user: ${JSON.stringify(returnObj)}`);
  return returnObj;
};
