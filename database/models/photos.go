package models

import (
	"time"
	"gorm.io/gorm"
)

type Photos struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Author      string         `json:"author"`
	ImageURL    string         `json:"image_url"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}