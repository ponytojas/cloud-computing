export const getValueCaseInsensitive = (obj, key) => {
  const lowercaseKey = key.toLowerCase();
  const keys = Object.keys(obj);

  for (const key of keys) {
    if (key.toLowerCase() === lowercaseKey) {
      return obj[key];
    }
  }

  return undefined;
};
