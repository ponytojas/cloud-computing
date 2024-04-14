import { createToken, deleteToken } from "../token/token.js";
import { logger } from "../utils/logger.js";

export const registerHandler = async (req, res) => {
  // try make request to process.env.DB_SERVICE_URL + "/users/create" with req.body and return the response
  try {
    const user = req.body;
    const result = await fetch(process.env.DB_SERVICE_URL + "/users/create", {
      method: "POST",
      body: JSON.stringify(user),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const body = await result.json();
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
    const result = await fetch(process.env.DB_SERVICE_URL + "/users/login", {
      method: "POST",
      body: JSON.stringify(user),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const body = await result.json();
    if (body.error) {
      res.status(401).json(body);
    } else {
      const token = await createToken(body.user);
      res.status(200).json({
        status: "OK",
        token,
        user: {
          userId: body.user.userId,
          username: body.user.username,
          email: body.user.email,
        },
      });
    }
  } catch (error) {
    console.error("Error logging in user:", error);
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
