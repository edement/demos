CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    username varchar(255) NOT NULL,
    email citext UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    is_trainer BOOLEAN NOT NULL DEFAULT FALSE
);