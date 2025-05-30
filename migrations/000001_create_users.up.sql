CREATE TABLE IF NOT EXISTS users (
                                     id UUID PRIMARY KEY,
                                     username TEXT UNIQUE NOT NULL,
                                     password TEXT NOT NULL,
                                     created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    is_active BOOLEAN NOT NULL DEFAULT true
    );
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_username ON users (username);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_id ON users (id);