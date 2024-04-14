import { logger } from "../utils/logger.js";

export const registerHandler = async (req, res) => {
  try {
    const result = await fetch(process.env.AUTH_SERVICE_URL + "/register", {
      method: "POST",
      body: JSON.stringify(req.body),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const body = await result.json();
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
    const result = await fetch(process.env.AUTH_SERVICE_URL + "/login", {
      method: "POST",
      body: JSON.stringify(req.body),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const body = await result.json();
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
    const result = await fetch(process.env.AUTH_SERVICE_URL + "/logout", {
      method: "POST",
      body: JSON.stringify(req.body),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const body = await result.json();
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
