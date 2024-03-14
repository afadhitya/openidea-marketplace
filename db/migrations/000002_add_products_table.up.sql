BEGIN;

DO
$$
BEGIN
IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'condition') THEN
    CREATE TYPE condition AS ENUM ('new', 'second');
  END IF;
END
$$;

CREATE TABLE IF NOT EXISTS products
(
    id UUID PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    price DECIMAL(13,2) NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    stock INTEGER NOT NULL,
    condition condition NOT NULL,
    tags TEXT NOT NULL,
    is_purchaseable BOOLEAN NOT NULL,
    purchase_count INTEGER NOT NULL,
    seller_id UUID NOT NULL REFERENCES users (id)
);

COMMIT;