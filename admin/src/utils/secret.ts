import CryptoJS from "crypto-js";

/**
 * @param word 待加密的字符
 * @param encryptSecretKey 加密秘钥
 * @param encryptIv 加密向量
 */
export const Encrypt = (word: string, encryptSecretKey: string, encryptIv: string): string => {
  const encryptSecretString = CryptoJS.enc.Utf8.parse(word);
  const encrypted = CryptoJS.AES.encrypt(encryptSecretString, CryptoJS.enc.Utf8.parse(encryptSecretKey), {
    iv: CryptoJS.enc.Utf8.parse(encryptIv),
    mode: CryptoJS.mode.CBC,
    padding: CryptoJS.pad.Pkcs7
  });
  return CryptoJS.enc.Base64.stringify(encrypted.ciphertext);
};

/**
 * @param word 待解密的字符
 * @param encryptSecretKey 加密秘钥
 * @param encryptIv 加密向量
 */
export const Decrypt = (word: string, encryptSecretKey: string, encryptIv: string): string => {
  const base64 = CryptoJS.enc.Base64.parse(word);
  const src = CryptoJS.enc.Base64.stringify(base64);
  const decrypt = CryptoJS.AES.decrypt(src, CryptoJS.enc.Utf8.parse(encryptSecretKey), {
    iv: CryptoJS.enc.Utf8.parse(encryptIv),
    mode: CryptoJS.mode.CBC,
    padding: CryptoJS.pad.Pkcs7
  });
  return CryptoJS.enc.Utf8.stringify(decrypt);
};
