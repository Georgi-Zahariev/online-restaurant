package managers

import (
	"errors"
	"time"

	"github.com/Georgi-Zahariev/online-restaurant/backend/models"
)

// Manager struct to manage Users, Dishes, and Orders
type Manager struct {
	Users  []models.User
	Dishes []models.Dish
	Orders []models.Order
}

// NewManager initializes a new Manager with hardcoded data
func NewManager() *Manager {
	return &Manager{
		Users: []models.User{

			{Base: models.Base{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now()}, PhoneNumber: "123456789"},
			{Base: models.Base{ID: "2", CreatedAt: time.Now(), UpdatedAt: time.Now()}, PhoneNumber: "987654321"},
		},
		Dishes: []models.Dish{
			{Base: models.Base{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "Pizza", Photo: "pizza.jpg", Price: 12.99, Description: "Delicious pizza", CategoryID: "1"},
			{Base: models.Base{ID: "2", CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "Burger", Photo: "burger.jpg", Price: 8.99, Description: "Juicy burger", CategoryID: "2"},
		},
		Orders: []models.Order{
			{Base: models.Base{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now()}, Price: 21.98, Status: "Pending", DayAndTime: time.Now(), UserID: "1"},
			{Base: models.Base{ID: "2", CreatedAt: time.Now(), UpdatedAt: time.Now()}, Price: 8.99, Status: "Completed", DayAndTime: time.Now(), UserID: "2"},
		},
	}
}

// GetAll returns all instances of a specific type
func (m *Manager) GetAll(entity string) (interface{}, error) {
	switch entity {
	case "users":
		return m.Users, nil
	case "dishes":
		return m.Dishes, nil
	case "orders":
		return m.Orders, nil
	default:
		return nil, errors.New("invalid entity type")
	}
}

// Get returns a specific instance by ID
func (m *Manager) Get(entity, id string) (interface{}, error) {
	switch entity {
	case "users":
		for _, user := range m.Users {
			if user.ID == id {
				return user, nil
			}
		}
	case "dishes":
		for _, dish := range m.Dishes {
			if dish.ID == id {
				return dish, nil
			}
		}
	case "orders":
		for _, order := range m.Orders {
			if order.ID == id {
				return order, nil
			}
		}
	default:
		return nil, errors.New("invalid entity type")
	}
	return nil, errors.New("not found")
}

// Create adds a new instance
func (m *Manager) Create(entity string, instance interface{}) error {
	switch entity {
	case "users":
		user, ok := instance.(models.User)
		if !ok {
			return errors.New("invalid instance type for users")
		}
		m.Users = append(m.Users, user)
	case "dishes":
		dish, ok := instance.(models.Dish)
		if !ok {
			return errors.New("invalid instance type for dishes")
		}
		m.Dishes = append(m.Dishes, dish)
	case "orders":
		order, ok := instance.(models.Order)
		if !ok {
			return errors.New("invalid instance type for orders")
		}
		m.Orders = append(m.Orders, order)
	default:
		return errors.New("invalid entity type")
	}
	return nil
}

// Update modifies an existing instance
func (m *Manager) Update(entity, id string, updatedInstance interface{}) error {
	switch entity {
	case "users":
		for i, user := range m.Users {
			if user.ID == id {
				m.Users[i] = updatedInstance.(models.User)
				return nil
			}
		}
	case "dishes":
		for i, dish := range m.Dishes {
			if dish.ID == id {
				m.Dishes[i] = updatedInstance.(models.Dish)
				return nil
			}
		}
	case "orders":
		for i, order := range m.Orders {
			if order.ID == id {
				m.Orders[i] = updatedInstance.(models.Order)
				return nil
			}
		}
	default:
		return errors.New("invalid entity type")
	}
	return errors.New("not found")
}

// Delete removes an instance by ID
func (m *Manager) Delete(entity, id string) error {
	switch entity {
	case "users":
		for i, user := range m.Users {
			if user.ID == id {
				m.Users = append(m.Users[:i], m.Users[i+1:]...)
				return nil
			}
		}
	case "dishes":
		for i, dish := range m.Dishes {
			if dish.ID == id {
				m.Dishes = append(m.Dishes[:i], m.Dishes[i+1:]...)
				return nil
			}
		}
	case "orders":
		for i, order := range m.Orders {
			if order.ID == id {
				m.Orders = append(m.Orders[:i], m.Orders[i+1:]...)
				return nil
			}
		}
	default:
		return errors.New("invalid entity type")
	}
	return errors.New("not found")
}
