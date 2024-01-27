import { defineStore } from "pinia";

export const useProductsStore = defineStore("products", () => {
  const products = ref([]);

  function updateProducts(prdts) {
    products.value = prdts;
  }

  return { products, updateProducts };
});
