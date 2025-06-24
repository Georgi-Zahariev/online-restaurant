package models

import "time"

// Base struct for common fields
type Base struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// User struct
type User struct {
	Base        `json:"base"`
	PhoneNumber string `json:"phone_number"`
}

// Address struct
type Address struct {
	Base    `json:"base"`
	City    string `json:"city"`
	Number  int    `json:"number"`
	Code    string `json:"code"`
	Country string `json:"country"`
	UserID  string `json:"user_id"`
}

// Category struct
type Category struct {
	Base `json:"base"`
	Name string `json:"name"`
}

// Dish struct
type Dish struct {
	Base        `json:"base"`
	Name        string  `json:"name"`
	Photo       string  `json:"photo"` // Not sure if it will work
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CategoryID  string  `json:"category_id"`
}

// Order struct
type Order struct {
	Base       `json:"base"`
	Price      float64   `json:"price"`
	Status     string    `json:"status"`
	DayAndTime time.Time `json:"day_and_time"`
	UserID     string    `json:"user_id"`
}

// OrderDish struct
type OrderItem struct {
	Base            `json:"base"`
	Count           int     `json:"count"`
	Price           float64 `json:"price"`
	Comments        string  `json:"comments"`
	CompletedByChef string  `json:"completed_by_chef"`
	OrderID         string  `json:"order_id"`
}

/*
// IDPUser struct
type IDPUser struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleID string `json:"role_id"`
}

// Role struct
type Role struct {
	ID       string `json:"id"`
	RoleName string `json:"role_name"`
}

// Privilage struct
type Privilege struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// RolePrivilege struct
type RolePrivilege struct {
	RoleID      string `json:"role_id"`
	PrivilegeID string `json:"privilege_id"`
}
*/
