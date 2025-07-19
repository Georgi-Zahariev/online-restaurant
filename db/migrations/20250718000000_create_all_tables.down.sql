-- Drop all tables in reverse order to respect foreign key constraints

DROP TABLE IF EXISTS "OrderItem" CASCADE;
DROP TABLE IF EXISTS "Order" CASCADE;
DROP TABLE IF EXISTS "Dish" CASCADE;
DROP TABLE IF EXISTS "Categories" CASCADE;
DROP TABLE IF EXISTS "Address" CASCADE;
DROP TABLE IF EXISTS "User" CASCADE;

-- Drop the UUID extension if no other objects depend on it
-- DROP EXTENSION IF EXISTS "pgcrypto";
