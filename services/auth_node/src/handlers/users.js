import { createToken, deleteToken } from "../token/token.js";
import { logger } from "../utils/logger.js";

import dotenv from "dotenv";
import axios from "axios";
dotenv.config();

export const registerHandler = async (req, res) => {
  try {
    const user = req.body;
    console.log("Register request:", JSON.stringify(user));
    const result = await axios.post(
      process.env.DB_SERVICE_URL + "/users/create",
      user
    );
    const body = await result.data;
    logger.debug(`User registered with ID: ${body.userId}`);
    res.status(201).json(body);
  } catch (error) {
    console.error("Error registering user:", error);
    res.status(500).send("ERROR 1001");
  }
};

export const loginUserHandler = async (req, res) => {
  const user = req.body;
  try {
    const result = await axios.post(
      process.env.DB_SERVICE_URL + "/users/login",
      user
    );
    const body = await result.data;
    console.log("Login response:", JSON.stringify(body));
    if (body.error) {
      res.status(401).json(body);
    } else {
      const token = await createToken(body.user);
      console.log(`User logged in: ${body.user.username}`);
      const resultObj = {
        status: "OK",
        token,
        user: {
          userId: body.user.userId,
          username: body.user.username,
          email: body.user.email,
        },
      };
      console.log("Login response:", resultObj);
      res.status(200).json(resultObj);
    }
  } catch (error) {
    console.log("Error logging in user:", error);
    res.status(500).send("ERROR 1002");
  }
};

export const logoutUserHandler = async (req, res) => {
  try {
    const { userId } = req.body;
    await deleteToken(userId);
    res.status(200).send("OK");
  } catch (error) {
    console.error("Error logging out user:", error);
    res.status(500).send("ERROR 1003");
  }
};
