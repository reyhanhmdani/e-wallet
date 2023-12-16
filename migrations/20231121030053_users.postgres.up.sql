CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    fullname VARCHAR(55),
    phone VARCHAR(55),
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(55),
    password VARCHAR(255) NOT NULL,
    email_verified_at TIMESTAMPTZ(0) DEFAULT NULL
    );
