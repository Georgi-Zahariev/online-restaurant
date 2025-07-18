package managers

import (
	"github.com/Georgi-Zahariev/online-restaurant/backend/models"
)

// GetAllOrder returns all orders from the database
func (m *Manager) GetAllOrder() ([]models.Order, error) {
	var orders []models.Order
	if err := m.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// GetOrder returns an order by ID
func (m *Manager) GetOrder(id string) (*models.Order, error) {
	var order models.Order
	if err := m.DB.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// CreateOrder adds a new order to the database
func (m *Manager) CreateOrder(order *models.Order) error {
	if err := m.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

// UpdateOrder updates an existing order by ID
func (m *Manager) UpdateOrder(id string, updated *models.Order) error {
	var order models.Order
	if err := m.DB.Where("id = ?", id).First(&order).Error; err != nil {
		return err
	}
	if err := m.DB.Model(&order).Updates(updated).Error; err != nil {
		return err
	}
	return nil
}

// DeleteOrder removes an order by ID
func (m *Manager) DeleteOrder(id string) error {
	if err := m.DB.Where("id = ?", id).Delete(&models.Order{}).Error; err != nil {
		return err
	}
	return nil
}
