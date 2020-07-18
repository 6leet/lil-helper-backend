package helpermodel

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UID      string `gorm:""`
	Username string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
	Admin    bool   `gorm:"default:false"`
	Active   bool   `gorm:"not null"`
}

type Mission struct {
	gorm.Model
	UID     string    `gorm:""`
	Content string    `gorm:"not null"`
	Picture string    `gorm:"not null"`
	Weight  []int     `gorm:"not null"`
	Score   int       `gorm:"not null"`
	Date    time.Time `gorm:"not null"`
	Active  bool      `gorm:"default:true"`
}

type Screenshot struct {
	gorm.Model
	UID       string    `gorm:""`
	UserID    string    `gorm:"not null"`
	MissionID string    `gorm:"not null"`
	Picture   string    `gorm:"not null"`
	Audit     bool      `gorm:"default:false"`
	Approve   bool      `gorm:"default:false"`
	Date      time.Time `gorm:"not null"`
}
