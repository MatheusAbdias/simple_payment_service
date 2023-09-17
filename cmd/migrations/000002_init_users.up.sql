CREATE TABLE
    IF NOT EXISTS users (
        id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
        email varchar(300) NOT NULL UNIQUE,
        full_name VARCHAR(300) NOT NULL,
        document varchar(14) NOT NULL UNIQUE
    );

CREATE INDEX IF NOT EXISTS users_email_idx ON users (email);

CREATE INDEX IF NOT EXISTS users_document_idx ON users (document);