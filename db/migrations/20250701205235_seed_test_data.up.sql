-- Seed User
INSERT INTO "User" (PhoneNumber) VALUES ('+359888123456');

-- Seed Address
INSERT INTO "Address" (City, Number, Code, Country, UserID)
VALUES ('Sofia', 12, '1000', 'Bulgaria', (SELECT ID FROM "User" LIMIT 1));

-- Seed Categories
INSERT INTO "Categories" (Name) VALUES ('Pizza'), ('Pasta'), ('Dessert');

-- Seed Dishes
INSERT INTO "Dish" (Name, Photo, Price, Description, CategoryID)
VALUES 
('Margherita', 'https://example.com/pizza.jpg', 9.99, 'Classic pizza with tomato and mozzarella', (SELECT ID FROM "Categories" WHERE Name = 'Pizza')),
('Spaghetti Carbonara', 'https://example.com/pasta.jpg', 12.50, 'Creamy pasta with bacon and egg', (SELECT ID FROM "Categories" WHERE Name = 'Pasta')),
('Tiramisu', 'https://example.com/dessert.jpg', 6.75, 'Coffee-flavored Italian dessert', (SELECT ID FROM "Categories" WHERE Name = 'Dessert'));

-- Seed Order
INSERT INTO "Order" (Price, Status, UserID)
VALUES (29.24, 'New', (SELECT ID FROM "User" LIMIT 1));

-- Seed Order Items
INSERT INTO "OrderItem" (Count, Price, Comments, CompletedByChef, OrderID)
VALUES 
(1, 9.99, 'No cheese', NULL, (SELECT ID FROM "Order" LIMIT 1)),
(1, 12.50, '', NULL, (SELECT ID FROM "Order" LIMIT 1)),
(1, 6.75, '', NULL, (SELECT ID FROM "Order" LIMIT 1));
