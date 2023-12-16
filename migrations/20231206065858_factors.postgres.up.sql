CREATE TABLE IF NOT EXISTS factors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    pin VARCHAR(100)
    );