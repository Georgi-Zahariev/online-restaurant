package models

import (
	"time"

	"gorm.io/gorm"
)

// Base struct for common fields matching your database schema
type Base struct {
	ID        string         `json:"id" gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:createdat"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updatedat"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"column:deletedat;index"`
}

// User struct
type User struct {
	Base
	PhoneNumber string `json:"phone_number" gorm:"column:phonenumber"`
}

// Address struct
type Address struct {
	Base
	City    string `json:"city" gorm:"column:city"`
	Number  int    `json:"number" gorm:"column:number"`
	Code    string `json:"code" gorm:"column:code"`
	Country string `json:"country" gorm:"column:country"`
	UserID  string `json:"user_id" gorm:"column:userid;type:uuid"`
}

// Category struct
type Category struct {
	Base
	Name string `json:"name" gorm:"column:name"`
}

// Dish struct
type Dish struct {
	Base
	Name        string  `json:"name" gorm:"column:name"`
	Photo       string  `json:"photo" gorm:"column:photo"`
	Price       float64 `json:"price" gorm:"column:price;type:numeric(10,2)"`
	Description string  `json:"description" gorm:"column:description"`
	CategoryID  string  `json:"category_id" gorm:"column:categoryid;type:uuid"`
}

// Order struct
type Order struct {
	Base
	Price      float64   `json:"price" gorm:"column:price;type:numeric(10,2)"`
	Status     string    `json:"status" gorm:"column:status"`
	DayAndTime time.Time `json:"day_and_time" gorm:"column:dayandtime"`
	UserID     string    `json:"user_id" gorm:"column:userid;type:uuid"`
}

// OrderItem struct
type OrderItem struct {
	Base
	Count           int     `json:"count" gorm:"column:count"`
	Price           float64 `json:"price" gorm:"column:price;type:numeric(10,2)"`
	Comments        string  `json:"comments" gorm:"column:comments"`
	CompletedByChef string  `json:"completed_by_chef" gorm:"column:completedbychef;type:uuid"`
	OrderID         string  `json:"order_id" gorm:"column:orderid;type:uuid"`
	DishID          string  `json:"dish_id" gorm:"column:dishid;type:uuid"`
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
