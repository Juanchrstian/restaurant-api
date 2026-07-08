CREATE TABLE sessions (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    status VARCHAR(20) NOT NULL,

    opened_by VARCHAR(100) NOT NULL,

    opening_cash BIGINT NOT NULL,

    closing_cash BIGINT,

    opened_at TIMESTAMP NOT NULL,

    closed_at TIMESTAMP,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    deleted_at TIMESTAMP

);

CREATE INDEX idx_sessions_deleted_at
ON sessions(deleted_at);