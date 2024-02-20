import { partial } from "filesize";

export const getFileSize = (value: number) => {
  const size = partial({ base: 2, standard: "jedec" });
  return size(value);
};
