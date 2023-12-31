CREATE TABLE IF NOT EXISTS topup (
    id VARCHAR(100) NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    amount BIGINT NOT NULL,
    status INT NOT NULL DEFAULT 0,
    snap_url VARCHAR(255)
    );