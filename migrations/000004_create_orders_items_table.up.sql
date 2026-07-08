CREATE TABLE order_items (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    order_id UUID NOT NULL,

    menu_id UUID NOT NULL,

    quantity INT NOT NULL,

    price BIGINT NOT NULL,

    subtotal BIGINT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_order_items_order
        FOREIGN KEY(order_id)
        REFERENCES orders(id),

    CONSTRAINT fk_order_items_menu
        FOREIGN KEY(menu_id)
        REFERENCES menus(id)

);

CREATE INDEX idx_order_items_order
ON order_items(order_id);

CREATE INDEX idx_order_items_menu
ON order_items(menu_id);