package managers

import (
	"context"

	"github.com/Georgi-Zahariev/online-restaurant/backend/middlewares"
	"gorm.io/gorm"
)

type Manager struct {
	DB *gorm.DB
}

func NewManager(db *gorm.DB) *Manager {
	return &Manager{DB: db}
}

// GetUserScopedDB returns a database connection scoped to a specific user
func (m *Manager) GetUserScopedDB(ctx context.Context) *gorm.DB {
	userID, ok := GetUserFromContext(ctx)
	if !ok {
		// Return unscoped DB if no user context (for admin operations)
		return m.DB
	}
	// Directly use GORM's Where method - simpler than custom scope
	return m.DB.Where("userid = ?", userID)
}

// GetUserFromContext extracts user ID from context using the middleware
func GetUserFromContext(ctx context.Context) (string, bool) {
	return middlewares.GetUserID(ctx)
}
