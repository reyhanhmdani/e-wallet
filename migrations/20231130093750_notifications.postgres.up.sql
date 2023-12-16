CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    status integer NOT NULL,
    title text NOT NULL,
    body text NOT NULL,
    is_read integer NOT NULL,
    created_at TIMESTAMP
);