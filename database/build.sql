-- Drop all tables
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS subscription_price;
DROP TABLE IF EXISTS item_price;
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

CREATE TYPE product_type AS ENUM ('subscription', 'item');

-- Create products table
CREATE TABLE products (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL NOT NULL,
    currency TEXT NOT NULL,
    stock INT NOT NULL,
    category TEXT,
    type product_type NOT NULL,
    attributes JSONB
);

-- Create subscription_price table
CREATE TABLE subscription_price (
    product_id UUID PRIMARY KEY REFERENCES products(id),
    monthly_price DECIMAL NOT NULL,  
);

-- Create item_price table
CREATE TABLE item_price (
    product_id UUID PRIMARY KEY REFERENCES products(id),
    one_time_price DECIMAL NOT NULL
);

-- Create carts table
CREATE TABLE carts (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    item_id UUID REFERENCES products(id),
    quantity INT DEFAULT 1,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create orders table
CREATE TABLE orders (
    user_id UUID REFERENCES users(id),
    product_id UUID REFERENCES products(id),
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (user_id, product_id, order_date)
);

CREATE TABLE active_subscriptions (
    user_id UUID REFERENCES users(id),
    product_id UUID REFERENCES products(id),
    start_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_date TIMESTAMP,
    adjusted_price DECIMAL,
    PRIMARY KEY (user_id, product_id)
);