package models

import "time"

// User struct
type User struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
}

// Address struct
type Address struct {
	ID      string `json:"id"`
	City    string `json:"city"`
	Number  int    `json:"number"`
	Code    string `json:"code"`
	Country string `json:"country"`
	UserID  string `json:"user_id"`
}

// Category struct
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Dish struct
type Dish struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Photo       string  `json:"photo"` // Assuming photo is a URL or base64 string
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CategoryID  string  `json:"category_id"`
}

// Order struct
type Order struct {
	ID         string    `json:"id"`
	Price      float64   `json:"price"`
	Status     string    `json:"status"`
	DayAndTime time.Time `json:"day_and_time"`
	UserID     string    `json:"user_id"`
}

// OrderDish struct
type OrderItem struct {
	ID              string  `json:"id"`
	Count           int     `json:"count"`
	Price           float64 `json:"price"`
	Comments        string  `json:"comments"`
	CompletedByChef string  `json:"completed_by_chef"` // Foreign key to User
	OrderID         string  `json:"order_id"`
}

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
