CREATE TABLE IF NOT EXISTS classes (
    id bigserial PRIMARY KEY,
    datetime timestamp(0) with time zone NOT NULL,
    location text NOT NULL,
    price numeric(10,2) DEFAULT 0.00,
    trainer_id bigint NOT NULL REFERENCES users(id),
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);