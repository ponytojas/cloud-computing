import axios from "axios";
import { logger } from "../utils/logger.js";

export const registerHandler = async (req, res) => {
  console.log("Register request:", req.body);
  try {
    const result = await axios.post(
      process.env.AUTH_SERVICE_URL + "/register",
      req.body
    );
    const body = await result.data;
    if (body.error) {
      res.status(400).json(body);
    } else {
      res.status(201).json(body);
    }
  } catch (error) {
    logger.error("Error registering user:", error);
    res.status(500).send("ERROR 1001");
  }
};

export const loginUserHandler = async (req, res) => {
  try {
    const result = await axios.post(
      process.env.AUTH_SERVICE_URL + "/login",
      req.body
    );
    const body = await result.data;
    logger.debug("Login response:", body);
    if (body.error) {
      res.status(400).json(body);
    } else {
      res.status(200).json({ ...body });
    }
  } catch (error) {
    logger.error("Error logging in user:", error);
    res.status(500).send("ERROR 1002");
  }
};

export const logoutUserHandler = async (req, res) => {
  try {
    const result = await axios.post(
      process.env.AUTH_SERVICE_URL + "/logout",
      req.body
    );
    const body = await result.data;
    if (body.error) {
      res.status(400).json(body);
    } else {
      await deleteToken(req.headers.authorization);
      res.status(200).json(body);
    }
  } catch (error) {
    logger.error("Error logging out user:", error);
    res.status(500).send("ERROR 1003");
  }
};
