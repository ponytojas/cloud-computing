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

  function clearBag() {
    items.value = [];
  }

  function add(item) {
    const existing = items.value.find((i) => i.Id === item.Id);
    if (existing) {
      existing.Added++;
    } else {
      items.value.push({ ...item, Added: 1 });
    }
  }

  function howManyInBag(productId) {
    const existing = items.value.find((i) => i.Id === productId);
    if (existing) {
      return existing.Added;
    } else {
      return 0;
    }
  }

  return { clearBag, add, howManyInBag, totalItems, items, total };
});
