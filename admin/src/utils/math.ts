import * as math from "mathjs";
// 加
export const add = (num1: number | string, num2: number | string): number => {
  const parsedNumber1 = typeof num1 === "string" ? parseFloat(num1) : num1;
  const parsedNumber2 = typeof num2 === "string" ? parseFloat(num2) : num2;
  return math.number(math.add(math.bignumber(parsedNumber1), math.bignumber(parsedNumber2)));
};

// 乘
export const multiply = (num1: number | string, num2: number | string): number => {
  const parsedNumber1 = typeof num1 === "string" ? parseFloat(num1) : num1;
  const parsedNumber2 = typeof num2 === "string" ? parseFloat(num2) : num2;
  return math.number(math.multiply(math.bignumber(parsedNumber1), math.bignumber(parsedNumber2)));
};

// 减
export const subtract = (num1: number | string, num2: number | string): number => {
  const parsedNumber1 = typeof num1 === "string" ? parseFloat(num1) : num1;
  const parsedNumber2 = typeof num2 === "string" ? parseFloat(num2) : num2;
  return math.number(math.subtract(math.bignumber(parsedNumber1), math.bignumber(parsedNumber2)));
};

// 除
export const divide = (num1: number | string, num2: number | string): number => {
  const parsedNumber1 = typeof num1 === "string" ? parseFloat(num1) : num1;
  const parsedNumber2 = typeof num2 === "string" ? parseFloat(num2) : num2;
  return math.number(math.divide(math.bignumber(parsedNumber1), math.bignumber(parsedNumber2)));
};
