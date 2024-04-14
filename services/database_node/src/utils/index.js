export const getValueCaseInsensitive = (obj, key) => {
  const lowercaseKey = key.toLowerCase();
  const keys = Object.keys(obj);

  for (let i = 0; i < keys.length; i++) {
    if (keys[i].toLowerCase() === lowercaseKey) {
      return obj[keys[i]];
    }
  }

  return undefined;
};
