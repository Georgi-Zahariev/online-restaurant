-- Drop all tables in reverse order to respect foreign key constraints

DROP TABLE IF EXISTS order_items CASCADE;
DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS dishes CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS addresses CASCADE;
DROP TABLE IF EXISTS users CASCADE;

-- Drop the UUID extension if no other objects depend on it
DROP EXTENSION IF EXISTS pgcrypto;
