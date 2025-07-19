package managers

import (
	"context"

	"github.com/Georgi-Zahariev/online-restaurant/backend/models"
)

// GetAllOrder returns all orders from the database for the current user
func (m *Manager) GetAllOrder(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	db := m.GetUserScopedDB(ctx)
	if err := db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// GetOrder returns an order by ID for the current user
func (m *Manager) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	var order models.Order
	db := m.GetUserScopedDB(ctx)
	if err := db.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// CreateOrder adds a new order to the database with user context
func (m *Manager) CreateOrder(ctx context.Context, order *models.Order) error {
	// Set the user ID from context
	if userID, ok := GetUserFromContext(ctx); ok {
		order.UserID = userID
	}

	if err := m.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

// UpdateOrder updates an existing order by ID for the current user
func (m *Manager) UpdateOrder(ctx context.Context, id string, updated *models.Order) error {
	var order models.Order
	db := m.GetUserScopedDB(ctx)
	if err := db.Where("id = ?", id).First(&order).Error; err != nil {
		return err
	}
	if err := db.Model(&order).Updates(updated).Error; err != nil {
		return err
	}
	return nil
}

// DeleteOrder removes an order by ID for the current user
func (m *Manager) DeleteOrder(ctx context.Context, id string) error {
	db := m.GetUserScopedDB(ctx)
	if err := db.Where("id = ?", id).Delete(&models.Order{}).Error; err != nil {
		return err
	}
	return nil
}
