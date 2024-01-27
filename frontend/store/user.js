import { defineStore } from "pinia";

export const useUserStore = defineStore("user", () => {
  const token = ref(null);
  // const isLogging = computed(() => true);
  const isLogging = computed(() => token?.value !== null);

  function login(token) {
    token.value = token;
  }

  function logout() {
    token.value = null;
  }

  return { token, isLogging, login, logout };
});
