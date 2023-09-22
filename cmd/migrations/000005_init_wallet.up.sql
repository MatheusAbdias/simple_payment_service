CREATE TABLE
    IF NOT EXISTS wallets (
        id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
        amount INT NOT NULL DEFAULT 0,
        owner_id uuid NOT NULL UNIQUE,
        FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
    );