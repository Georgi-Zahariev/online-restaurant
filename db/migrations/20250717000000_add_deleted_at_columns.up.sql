-- Add DeletedAt column to existing tables for soft delete support

-- Add DeletedAt to User table
ALTER TABLE "User" ADD COLUMN DeletedAt TIMESTAMP;
CREATE INDEX idx_user_deleted_at ON "User"(DeletedAt);

-- Add DeletedAt to Address table
ALTER TABLE "Address" ADD COLUMN DeletedAt TIMESTAMP;
CREATE INDEX idx_address_deleted_at ON "Address"(DeletedAt);

-- Add DeletedAt to Categories table
ALTER TABLE "Categories" ADD COLUMN DeletedAt TIMESTAMP;
CREATE INDEX idx_categories_deleted_at ON "Categories"(DeletedAt);

-- Add DeletedAt to Dish table
ALTER TABLE "Dish" ADD COLUMN DeletedAt TIMESTAMP;
CREATE INDEX idx_dish_deleted_at ON "Dish"(DeletedAt);

-- Add DeletedAt to Order table
ALTER TABLE "Order" ADD COLUMN DeletedAt TIMESTAMP;
CREATE INDEX idx_order_deleted_at ON "Order"(DeletedAt);

-- Add DeletedAt to OrderItem table
ALTER TABLE "OrderItem" ADD COLUMN DeletedAt TIMESTAMP;
CREATE INDEX idx_orderitem_deleted_at ON "OrderItem"(DeletedAt);
