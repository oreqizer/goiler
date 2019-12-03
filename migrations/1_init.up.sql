/*
=== ADD-ONS
 */
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

/*
=== ACCOUNT ===
 */

CREATE TABLE account
(
    id         UUID PRIMARY KEY     DEFAULT uuid_generate_v1mc(),
    auth_id    TEXT UNIQUE NOT NULL,
    name       TEXT        NOT NULL,
    surname    TEXT        NOT NULL,
    email      TEXT        NOT NULL,
    is_admin   BOOL        NOT NULL DEFAULT FALSE,
    -- misc
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX account_auth_id ON account (auth_id);
