import { logger } from "./logger.js";

export const getValueCaseInsensitive = (obj, key) => {
  const lowercaseKey = key.toLowerCase();
  const keys = Object.keys(obj);

  logger.info("KEYS: " + keys);
  logger.info("VALUES: " + JSON.stringify(obj));

  for (const key of keys) {
    if (key.toLowerCase() === lowercaseKey) {
      return obj[key];
    }
  }

  return undefined;
};
