-- Drop all tables
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS orders;

-- Create users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    cart_id UUID REFERENCES carts(id)
);

CREATE TYPE payment_type AS ENUM ('onetime', 'recurring');
CREATE TYPE category AS ENUM ('smartphone', 'tablet', 'accessory', 'wearable', 'subscription');

-- Create products table
CREATE TABLE products (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL NOT NULL,
    currency TEXT NOT NULL,
    stock INT NOT NULL,
    category category NOT NULL,
    payment_type payment_type NOT NULL,
    attributes JSONB
);

-- Create carts table
CREATE TABLE carts (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    item_id INTEGER REFERENCES products(id),
    quantity INT DEFAULT 1,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create orders table
CREATE TABLE orders (
    user_id UUID REFERENCES users(id),
    product_id INTEGER REFERENCES products(id),
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (user_id, product_id, order_date)
);

CREATE TABLE active_subscriptions (
    user_id UUID REFERENCES users(id),
    product_id INTEGER REFERENCES products(id),
    start_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_date TIMESTAMP,
    adjusted_price DECIMAL,
    PRIMARY KEY (user_id, product_id)
);