import { defineStore } from "pinia";

export const useUserStore = defineStore("user", () => {
  const token = ref(null);
  const user = ref(null);

  const isLogging = computed(() => token?.value !== null);

  function login(data) {
    token.value = data.token;
    user.value = data.user;
  }

  function logout() {
    token.value = null;
    user.value = null;
  }

  return { token, user, isLogging, login, logout };
});
