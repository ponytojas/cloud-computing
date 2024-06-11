CREATE TABLE IF NOT EXISTS "users" (
  user_id SERIAL PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "auth" (
  auth_id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES "users"(user_id) ON DELETE CASCADE,
  password_hash CHAR(60) NOT NULL,
  last_login TIMESTAMP WITH TIME ZONE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "product" (
  product_id SERIAL PRIMARY KEY,
  name VARCHAR(50) UNIQUE NOT NULL,
  pricing NUMERIC(10, 2) NOT NULL,
  description VARCHAR(100) NOT NULL,
  rating NUMERIC(2, 1) CHECK (
    rating >= 0
    AND rating <= 5
  ),
  picture TEXT,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "product_stock" (
  product_stock_id SERIAL PRIMARY KEY,
  product_id INTEGER REFERENCES "product"(product_id) ON DELETE CASCADE,
  quantity INTEGER NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT unique_product_id UNIQUE (product_id)
);

CREATE TABLE IF NOT EXISTS "user_purchases" (
  purchase_id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES "users"(user_id) ON DELETE CASCADE,
  purchase_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "purchase_items" (
  item_id SERIAL PRIMARY KEY,
  purchase_id INTEGER REFERENCES "user_purchases"(purchase_id) ON DELETE CASCADE,
  product_id INTEGER REFERENCES "product"(product_id) ON DELETE CASCADE,
  quantity INTEGER NOT NULL,
  price_per_unit NUMERIC(10, 2) NOT NULL,
  total_price NUMERIC(10, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS "invoices" (
  invoice_id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES "users"(user_id) ON DELETE CASCADE,
  purchase_id INTEGER REFERENCES "user_purchases"(purchase_id) ON DELETE CASCADE,
  total_amount NUMERIC(10, 2) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);