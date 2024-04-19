import { create } from "zustand";

export const useStore = create((set) => ({
  token: null,
  userId: null,
  cart: {},
  setToken: (token) => set({ token }),
  setUserId: (userId) => set({ userId }),
  setCart: (cart) => set({ cart }),
}));
