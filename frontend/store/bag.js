import { defineStore } from "pinia";

export const useBagStore = defineStore("bag", () => {
  const items = ref([]);
  const total = computed(() => {
    let value = 0;
    for (const item of items.value) {
      value += item.Pricing * item.Added;
    }
    return value;
  });

  const totalItems = computed(() => {
    let value = 0;
    for (const item of items.value) {
      value += item.Added;
    }
    return value;
  });

  function clearBag(token) {
    token.value = [];
  }

  return { clearBag, totalItems, items, total };
});
