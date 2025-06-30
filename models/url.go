package models

import (
	"time"
)

type URL struct {
	ID          uint   `gorm:"primaryKey"`
	ShortCode   string `gorm:"uniqueIndex"`
	OriginalURL string
	ClickCount  uint
	CreatedAt   time.Time
}
