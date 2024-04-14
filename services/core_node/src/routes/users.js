import { Router } from "express";
import {
  loginUserHandler,
  logoutUserHandler,
  registerHandler,
} from "../handlers/users.js";

export const getUsersRoute = () => {
  const userRouter = Router();
  userRouter.post("/register", registerHandler);
  userRouter.post("/login", loginUserHandler);
  userRouter.post("/logout", logoutUserHandler);
  return userRouter;
};
