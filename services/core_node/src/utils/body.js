const KEYS = ["name", "pricing", "description"];
const RES_KEYS = [
  "name",
  "pricing",
  "description",
  "product_id",
  "product_stock_id",
  "quantity",
];

export const adaptBody = (data) => {
  const newData = data.map((d) => {
    const _body = {};
    RES_KEYS.forEach((key) => {
      if (d[key]) {
        _body[key] = d[key];
      }
    });

    // Check if _body has all the required keys
    if (Object.keys(_body).length !== RES_KEYS.length) {
      return null;
    }
    return _body;
  });
  return newData;
};

export const adaptResponse = (body) => {
  const _body = {};
  RES_KEYS.forEach((key) => {
    if (body[key]) {
      _body[key] = body[key];
    }
  });

  // Check if _body has all the required keys
  if (Object.keys(_body).length !== RES_KEYS.length) {
    return null;
  }
  return _body;
};
