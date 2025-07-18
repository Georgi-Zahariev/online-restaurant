package managers

import "gorm.io/gorm"

type Manager struct {
	DB *gorm.DB
}

func NewManager(db *gorm.DB) *Manager {
	return &Manager{DB: db}
}
