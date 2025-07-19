package managers

import (
	"context"
	"fmt"

	"github.com/Georgi-Zahariev/online-restaurant/backend/models"
)

// GetAllUsers returns all users from the database (admin only operation)
func (m *Manager) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	// For users, we typically don't scope by user unless it's for admin operations
	if err := m.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUser returns a user by ID
func (m *Manager) GetUser(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	if err := m.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetCurrentUser returns the current user from context
func (m *Manager) GetCurrentUser(ctx context.Context) (*models.User, error) {
	userID, ok := GetUserFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("no user in context")
	}
	return m.GetUser(ctx, userID)
}

// CreateUser adds a new user to the database
func (m *Manager) CreateUser(ctx context.Context, user *models.User) error {
	if err := m.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser updates an existing user by ID
func (m *Manager) UpdateUser(ctx context.Context, id string, updated *models.User) error {
	var user models.User
	if err := m.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	if err := m.DB.Model(&user).Updates(updated).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser removes a user by ID
func (m *Manager) DeleteUser(ctx context.Context, id string) error {
	if err := m.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}
