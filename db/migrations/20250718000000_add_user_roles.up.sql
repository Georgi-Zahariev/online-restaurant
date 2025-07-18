-- Add role column to User table
ALTER TABLE "User" ADD COLUMN role VARCHAR(20) DEFAULT 'customer' CHECK (role IN ('customer', 'kitchen', 'delivery', 'owner'));

-- Add timestamps columns that might be missing
ALTER TABLE "User" ADD COLUMN IF NOT EXISTS CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE "User" ADD COLUMN IF NOT EXISTS UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- Update existing users to have customer role by default
UPDATE "User" SET role = 'customer' WHERE role IS NULL;
