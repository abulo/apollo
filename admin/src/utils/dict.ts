import { useDictStore } from "@/stores/modules/dict";
export const getDict = (label: string) => {
  const dictStore = useDictStore();
  const dictList = dictStore.dictListGet[label] || [];
  return dictList;
};
