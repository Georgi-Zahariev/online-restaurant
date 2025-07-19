package managers

import (
	"context"

	"github.com/Georgi-Zahariev/online-restaurant/backend/models"
)

// GetAllDishes returns all dishes from the database (global, not user-scoped)
func (m *Manager) GetAllDishes(ctx context.Context) ([]models.Dish, error) {
	var dishes []models.Dish
	if err := m.DB.Find(&dishes).Error; err != nil {
		return nil, err
	}
	return dishes, nil
}

// GetDish returns a dish by ID (global, not user-scoped)
func (m *Manager) GetDish(ctx context.Context, id string) (*models.Dish, error) {
	var dish models.Dish
	if err := m.DB.Where("id = ?", id).First(&dish).Error; err != nil {
		return nil, err
	}
	return &dish, nil
}

// CreateDish adds a new dish to the database (admin operation)
func (m *Manager) CreateDish(ctx context.Context, dish *models.Dish) error {
	if err := m.DB.Create(dish).Error; err != nil {
		return err
	}
	return nil
}

// UpdateDish updates an existing dish by ID (admin operation)
func (m *Manager) UpdateDish(ctx context.Context, id string, updated *models.Dish) error {
	var dish models.Dish
	if err := m.DB.Where("id = ?", id).First(&dish).Error; err != nil {
		return err
	}
	if err := m.DB.Model(&dish).Updates(updated).Error; err != nil {
		return err
	}
	return nil
}

// DeleteDish removes a dish by ID (admin operation)
func (m *Manager) DeleteDish(ctx context.Context, id string) error {
	if err := m.DB.Where("id = ?", id).Delete(&models.Dish{}).Error; err != nil {
		return err
	}
	return nil
}
