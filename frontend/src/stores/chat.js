import { defineStore } from "pinia";

export const useChatStore = defineStore({
  id: "chat",
  state: () => ({
    lastDiv: null,
  }),
  getters: {
    GET_LAST_DIV: (state) => state.lastDiv,
  },
  actions: {},
});
