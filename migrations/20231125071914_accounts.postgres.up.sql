CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    account_number VARCHAR(100),
    balance NUMERIC(19, 2)
    );
