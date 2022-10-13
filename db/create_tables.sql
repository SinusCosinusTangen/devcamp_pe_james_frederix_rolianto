CREATE TABLE IF NOT EXISTS product (
    id bigserial PRIMARY KEY,
    name varchar NOT NULL,
    image_url text NOT NULL,
    description text NOT NULL,
    weight int NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS variant (
    id bigserial PRIMARY KEY,
    product_id int NOT NULL,
    name varchar NOT NULL,
    image_url text,
    description text NOT NULL,
    weight int NOT NULL,
    price int NOT NULL,
    stock int NOT NULL,
    discount int NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);