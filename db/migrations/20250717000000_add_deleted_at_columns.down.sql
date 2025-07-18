-- Remove DeletedAt columns (rollback migration)

-- Remove DeletedAt from User table
DROP INDEX IF EXISTS idx_user_deleted_at;
ALTER TABLE "User" DROP COLUMN IF EXISTS DeletedAt;

-- Remove DeletedAt from Address table
DROP INDEX IF EXISTS idx_address_deleted_at;
ALTER TABLE "Address" DROP COLUMN IF EXISTS DeletedAt;

-- Remove DeletedAt from Categories table
DROP INDEX IF EXISTS idx_categories_deleted_at;
ALTER TABLE "Categories" DROP COLUMN IF EXISTS DeletedAt;

-- Remove DeletedAt from Dish table
DROP INDEX IF EXISTS idx_dish_deleted_at;
ALTER TABLE "Dish" DROP COLUMN IF EXISTS DeletedAt;

-- Remove DeletedAt from Order table
DROP INDEX IF EXISTS idx_order_deleted_at;
ALTER TABLE "Order" DROP COLUMN IF EXISTS DeletedAt;

-- Remove DeletedAt from OrderItem table
DROP INDEX IF EXISTS idx_orderitem_deleted_at;
ALTER TABLE "OrderItem" DROP COLUMN IF EXISTS DeletedAt;
