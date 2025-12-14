CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  category_name VARCHAR(200) NOT NULL,
  description TEXT
);

CREATE TABLE items (
  id SERIAL PRIMARY KEY,
  category_id INT,
  item_name VARCHAR(200) NOT NULL,
  quantity INT NOT NULL DEFAULT 1,
  price NUMERIC(12,2),
  purchase_date DATE,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),

  CONSTRAINT fk_category FOREIGN KEY (category_id)
    REFERENCES categories(id)
);
