-- Base table for common fields
CREATE TABLE Base (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- User table
CREATE TABLE User (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    PhoneNumber VARCHAR(50) NOT NULL
);

-- Address table
CREATE TABLE Address (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    City VARCHAR(50) NOT NULL,
    Number INT NOT NULL,
    Code VARCHAR(50) NOT NULL,
    Country VARCHAR(50) NOT NULL,
    UserID UUID NOT NULL REFERENCES User(ID) ON DELETE CASCADE
);

-- Categories table
CREATE TABLE Categories (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Name VARCHAR(50) NOT NULL
);

-- Dish table
CREATE TABLE Dish (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Name VARCHAR(100) NOT NULL,
    Photo BYTEA, -- Storing binary data for photos
    Price NUMERIC(10, 2) NOT NULL,
    Description VARCHAR(500),
    CategoryID UUID NOT NULL REFERENCES Categories(ID) ON DELETE CASCADE
);

-- Order table
CREATE TABLE "Order" (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Price NUMERIC(10, 2) NOT NULL,
    Status VARCHAR(50) NOT NULL CHECK (Status IN ('Pending', 'Completed', 'Cancelled')), -- Limited to specific options
    DayAndTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UserID UUID NOT NULL REFERENCES User(ID) ON DELETE CASCADE
);

-- OrderItem table
CREATE TABLE OrderItem (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Count INT NOT NULL,
    Price NUMERIC(10, 2) NOT NULL,
    Comments VARCHAR(500),
    CompletedByChef UUID REFERENCES User(ID), -- Foreign key to User
    OrderID UUID NOT NULL REFERENCES "Order"(ID) ON DELETE CASCADE
);

-- IDPUser table
CREATE TABLE IDPUser (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Name VARCHAR(50) NOT NULL,
    Email VARCHAR(50) NOT NULL UNIQUE,
    RoleID UUID NOT NULL REFERENCES Role(ID) ON DELETE CASCADE
);

-- Role table
CREATE TABLE Role (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    RoleName VARCHAR(50) NOT NULL
);

-- Privilage table
CREATE TABLE Privilage (
    PrivilageID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    PrivilageName VARCHAR(50) NOT NULL
);

-- RolePrivilages table
CREATE TABLE RolePrivilages (
    RoleID UUID NOT NULL REFERENCES Role(ID) ON DELETE CASCADE,
    PrivilageID UUID NOT NULL REFERENCES Privilage(PrivilageID) ON DELETE CASCADE,
    PRIMARY KEY (RoleID, PrivilageID)
);