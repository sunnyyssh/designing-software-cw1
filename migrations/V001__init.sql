CREATE TABLE bank_accounts (
    id      UUID PRIMARY KEY,
    name    TEXT NOT NULL,
    balance BIGINT NOT NULL,
    blocked BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE categories (
    id   UUID PRIMARY KEY,
    type TEXT NOT NULL,
    name TEXT NOT NULL
);

CREATE TABLE operations (
    id           UUID PRIMARY KEY,
    account_id   UUID NOT NULL,
    type         TEXT NOT NULL,
    amount       BIGINT NOT NULL,
    time         TIMESTAMP NOT NULL,
    description  TEXT NOT NULL,
    category_id  UUID REFERENCES categories (id) ON DELETE SET NULL,
);