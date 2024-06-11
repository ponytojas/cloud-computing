import { create } from "zustand";

export const useStore = create((set) => ({
  token: null,
  userId: null,
  cart: {},
  products: [],
  setToken: (token) => set({ token }),
  setUserId: (userId) => set({ userId }),
  setCart: (cart) => set({ cart }),
  setProducts: (products) => set({ products }),
}));
