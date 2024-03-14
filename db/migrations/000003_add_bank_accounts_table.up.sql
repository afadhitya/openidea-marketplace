BEGIN;

CREATE TABLE IF NOT EXISTS bank_accounts 
(
    id UUID PRIMARY KEY,
    bank_name VARCHAR(255) NOT NULL,
    bank_account_name VARCHAR(255) NOT NULL,
    bank_account_number VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL REFERENCES users (id)
);

COMMIT;