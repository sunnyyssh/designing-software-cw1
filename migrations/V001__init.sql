CREATE TABLE bank_accounts (
    id      UUID PRIMARY KEY,
    name    VARCHAR(255) NOT NULL,
    balance BIGINT NOT NULL,
    blocked BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE categories (
    id   UUID PRIMARY KEY,
    type VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE operations (
    id           UUID PRIMARY KEY,
    account_id   UUID NOT NULL,
    type         VARCHAR(255) NOT NULL,
    amount       BIGINT NOT NULL,
    time         TIMESTAMP NOT NULL,
    description  TEXT NOT NULL,
    category_id  UUID REFERENCES categories (id) ON DELETE SET NULL
);
