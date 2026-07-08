CREATE TABLE orders (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    session_id UUID NOT NULL,

    order_number VARCHAR(50) NOT NULL UNIQUE,

    status VARCHAR(20) NOT NULL,

    total_amount BIGINT NOT NULL DEFAULT 0,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    deleted_at TIMESTAMP,

    CONSTRAINT fk_orders_session
        FOREIGN KEY(session_id)
        REFERENCES sessions(id)

);

CREATE INDEX idx_orders_session
ON orders(session_id);

CREATE INDEX idx_orders_deleted
ON orders(deleted_at);