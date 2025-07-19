-- Enable UUID generation
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- User table
CREATE TABLE "User" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deletedat TIMESTAMP,
    phonenumber VARCHAR(50) NOT NULL
);

-- Create index for soft delete
CREATE INDEX idx_user_deletedat ON "User"(deletedat);

-- Address table
CREATE TABLE "Address" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deletedat TIMESTAMP,
    city VARCHAR(50) NOT NULL,
    number INT NOT NULL,
    code VARCHAR(50) NOT NULL,
    country VARCHAR(50) NOT NULL,
    userid UUID NOT NULL REFERENCES "User"(id) ON DELETE CASCADE
);

-- Create index for soft delete
CREATE INDEX idx_address_deletedat ON "Address"(deletedat);

-- Categories table
CREATE TABLE "Categories" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deletedat TIMESTAMP,
    name VARCHAR(50) NOT NULL
);

-- Create index for soft delete
CREATE INDEX idx_categories_deletedat ON "Categories"(deletedat);

-- Dish table
CREATE TABLE "Dish" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deletedat TIMESTAMP,
    name VARCHAR(100) NOT NULL,
    photo VARCHAR(1024),
    price NUMERIC(10, 2) NOT NULL,
    description VARCHAR(500),
    categoryid UUID NOT NULL REFERENCES "Categories"(id) ON DELETE CASCADE
);

-- Create index for soft delete
CREATE INDEX idx_dish_deletedat ON "Dish"(deletedat);

-- Order table
CREATE TABLE "Order" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deletedat TIMESTAMP,
    price NUMERIC(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('Draft', 'New', 'In Progress', 'Completed', 'Preparing', 'Delivered')),
    dayandtime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    userid UUID NOT NULL REFERENCES "User"(id) ON DELETE CASCADE
);

-- Create index for soft delete
CREATE INDEX idx_order_deletedat ON "Order"(deletedat);

-- OrderItem table
CREATE TABLE "OrderItem" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deletedat TIMESTAMP,
    count INT NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    comments VARCHAR(500),
    completedbychef UUID REFERENCES "User"(id),
    orderid UUID NOT NULL REFERENCES "Order"(id) ON DELETE CASCADE
);

-- Create index for soft delete
CREATE INDEX idx_orderitem_deletedat ON "OrderItem"(deletedat);

-- Insert sample categories
INSERT INTO "Categories" (name) VALUES 
('Pizza'), 
('Pasta'), 
('Dessert'),
('Appetizers'),
('Main Course'),
('Beverages');

-- Insert sample dishes
INSERT INTO "Dish" (name, photo, price, description, categoryid)
VALUES 
('Margherita Pizza', 'https://example.com/pizza.jpg', 9.99, 'Classic pizza with tomato and mozzarella', (SELECT id FROM "Categories" WHERE name = 'Pizza')),
('Spaghetti Carbonara', 'https://example.com/pasta.jpg', 12.50, 'Creamy pasta with bacon and egg', (SELECT id FROM "Categories" WHERE name = 'Pasta')),
('Tiramisu', 'https://example.com/dessert.jpg', 6.75, 'Coffee-flavored Italian dessert', (SELECT id FROM "Categories" WHERE name = 'Dessert')),
('Caesar Salad', 'https://example.com/salad.jpg', 8.50, 'Fresh romaine lettuce with caesar dressing', (SELECT id FROM "Categories" WHERE name = 'Appetizers')),
('Grilled Chicken', 'https://example.com/chicken.jpg', 15.99, 'Juicy grilled chicken breast with herbs', (SELECT id FROM "Categories" WHERE name = 'Main Course')),
('Coca Cola', 'https://example.com/coke.jpg', 2.50, 'Classic soft drink', (SELECT id FROM "Categories" WHERE name = 'Beverages'));

-- Insert test users with specific UUIDs for testing user-scoped functionality
INSERT INTO "User" (id, phonenumber) VALUES 
('550e8400-e29b-41d4-a716-446655440001', '+359888123456'),
('550e8400-e29b-41d4-a716-446655440002', '+359888654321'),
('550e8400-e29b-41d4-a716-446655440003', '+359888999888');

-- Insert sample addresses for the test users
INSERT INTO "Address" (city, number, code, country, userid)
VALUES 
('Sofia', 12, '1000', 'Bulgaria', '550e8400-e29b-41d4-a716-446655440001'),
('Plovdiv', 25, '4000', 'Bulgaria', '550e8400-e29b-41d4-a716-446655440002'),
('Varna', 8, '9000', 'Bulgaria', '550e8400-e29b-41d4-a716-446655440003');

-- Insert sample orders for each user to test user-scoped functionality
INSERT INTO "Order" (price, status, userid) VALUES 
(19.99, 'Preparing', '550e8400-e29b-41d4-a716-446655440001'),
(25.50, 'Delivered', '550e8400-e29b-41d4-a716-446655440001'),
(12.75, 'New', '550e8400-e29b-41d4-a716-446655440002'),
(35.20, 'Preparing', '550e8400-e29b-41d4-a716-446655440002'),
(8.99, 'New', '550e8400-e29b-41d4-a716-446655440003');

-- Insert sample order items
INSERT INTO "OrderItem" (count, price, comments, orderid)
VALUES 
(1, 9.99, 'No cheese please', (SELECT id FROM "Order" WHERE userid = '550e8400-e29b-41d4-a716-446655440001' AND price = 19.99 LIMIT 1)),
(1, 12.50, '', (SELECT id FROM "Order" WHERE userid = '550e8400-e29b-41d4-a716-446655440001' AND price = 25.50 LIMIT 1)),
(2, 6.75, 'Extra sweet', (SELECT id FROM "Order" WHERE userid = '550e8400-e29b-41d4-a716-446655440002' AND price = 12.75 LIMIT 1)),
(1, 15.99, 'Well done', (SELECT id FROM "Order" WHERE userid = '550e8400-e29b-41d4-a716-446655440002' AND price = 35.20 LIMIT 1)),
(1, 8.99, '', (SELECT id FROM "Order" WHERE userid = '550e8400-e29b-41d4-a716-446655440003' LIMIT 1));
