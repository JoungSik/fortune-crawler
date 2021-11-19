package models

import "time"

type Fortune struct {
	Name      string
	Content   string
	DueDate   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
