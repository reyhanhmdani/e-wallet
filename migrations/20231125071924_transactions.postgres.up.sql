CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sof_number VARCHAR(100),
    dof_number VARCHAR(100),
    amount NUMERIC(19, 2),
    transaction_type VARCHAR(1),
    account_id UUID REFERENCES accounts (id) ON DELETE CASCADE,
    transactions_datetime TIMESTAMP
    );
