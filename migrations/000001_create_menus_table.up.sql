CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE menus (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name VARCHAR(100) NOT NULL,

    description TEXT,

    price BIGINT NOT NULL,

    stock INTEGER NOT NULL DEFAULT 0,

    available BOOLEAN NOT NULL DEFAULT TRUE,

    image_url TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    deleted_at TIMESTAMP NULL

);