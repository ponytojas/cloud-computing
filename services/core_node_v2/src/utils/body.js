const RES_KEYS = [
  "name",
  "pricing",
  "description",
  "product_id",
  "product_stock_id",
  "rating",
  "picture",
  "quantity",
];

export const adaptResponse = (body) => {
  const _body = {};
  RES_KEYS.forEach((key) => {
    if (body[key]) {
      _body[key] = body[key];
    }
  });

  return _body;
};
