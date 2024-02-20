import { defineStore } from "pinia";
import { DictState } from "@/stores/interface";
import piniaPersistConfig from "@/stores/helper/persist";
// import { getDictListApi } from "@/api/modules/systemDictLabel";
// import { SystemDict } from "@/api/interface/systemDict";

export const useDictStore = defineStore({
  id: "geeker-dict",
  state: (): DictState => ({
    dict: {}
  }),
  getters: {
    dictListGet: state => state.dict
  },
  actions: {
    // async getDictList() {
    //   const { data } = await getDictListApi();
    //   this.dict = data;
    // }
  },
  persist: piniaPersistConfig("geeker-dict")
});
