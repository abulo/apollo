import { multiply, divide } from "@/utils/math";

/**
 * 元转分
 */
export const yuanToFen = (price: string | number): number => {
  return multiply(price, 100);
};

/**
 * 分转元
 */
export const fenToYuan = (price: string | number): number => {
  return Number(divide(price, 100).toFixed(2));
};
