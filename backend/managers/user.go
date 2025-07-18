package managers

import (
	"github.com/Georgi-Zahariev/online-restaurant/backend/models"
)

// GetAllUsers returns all users from the database
func (m *Manager) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := m.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUser returns a user by ID
func (m *Manager) GetUser(id string) (*models.User, error) {
	var user models.User
	if err := m.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser adds a new user to the database
func (m *Manager) CreateUser(user *models.User) error {
	if err := m.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser updates an existing user by ID
func (m *Manager) UpdateUser(id string, updated *models.User) error {
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
func (m *Manager) DeleteUser(id string) error {
	if err := m.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}
