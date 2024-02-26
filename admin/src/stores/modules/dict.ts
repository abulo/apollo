import { defineStore } from "pinia";
import { DictState } from "@/stores/interface";
import piniaPersistConfig from "@/stores/helper/persist";
import { getAllSystemDictApi } from "@/api/modules/systemDict";

export const useDictStore = defineStore({
  id: "geeker-dict",
  state: (): DictState => ({
    dict: {}
  }),
  getters: {
    dictListGet: state => state.dict
  },
  actions: {
    async getDictList() {
      const { data } = await getAllSystemDictApi();
      this.dict = data;
    }
  },
  persist: piniaPersistConfig("geeker-dict")
});
