import { createUser, loginUser } from "../database/users.js";
import { getValueCaseInsensitive } from "../utils/index.js";
import { logger } from "../utils/logger.js";

export const createUserHandler = async (req, res) => {
  try {
    const user = req.body;
    const createdUserId = await createUser(user);
    logger.debug(`User created with ID: ${createdUserId}`);
    res.status(200).json({ userId: createdUserId });
  } catch (error) {
    console.error("Error creating user:", error);
    res.status(500).send("ERROR 1001");
  }
};

export const loginUserHandler = async (req, res) => {
  try {
    const username = getValueCaseInsensitive(req.body, "username");
    const password = getValueCaseInsensitive(req.body, "password");
    logger.info(
      `User ${username} is trying to log in with password ${password}`
    );
    const authCheck = await loginUser(username, password);
    logger.debug(`User logged in with ID: ${authCheck.userId}`);
    res.status(200).json({ user: authCheck });
  } catch (error) {
    console.error("Error logging in user:", error);
    res.status(500).send("ERROR 1002");
  }
};
