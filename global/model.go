package global

import (
	"gorm.io/gorm"
	"time"
)

type GvaModel struct {
	ID        uint `gorm:"primary_key;auto_increment"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}