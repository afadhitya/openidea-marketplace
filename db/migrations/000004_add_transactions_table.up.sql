BEGIN;

CREATE TABLE IF NOT EXISTS transactions 
(
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL REFERENCES products (id),
    quantity INTEGER NOT NULL,
    payment_proof_image_url VARCHAR(255) NOT NULL,
    buyer_id UUID NOT NULL REFERENCES users (id)
);

COMMIT;