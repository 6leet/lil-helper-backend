package helpermodel

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UID         string `gorm:""`
	Username    string `gorm:"unique_index;not null"`
	Password    string `gorm:"not null"`
	Nickname    string `gorm:"default:'__nickname__'"`
	Admin       bool   `gorm:"default:false"`
	Active      bool   `gorm:"not null;default:true"`
	Score       int    `gorm:"default:0"`
	Exp         int    `gorm:"default:0"`
	Level       uint   `gorm:"default:0"`
	Email       string `gorm:"not null"`
	Certificate bool   `gorm:"default:false"`
}

type Mission struct {
	gorm.Model
	UID        string    `gorm:""`
	Content    string    `gorm:"not null"`
	Picture    string    `gorm:"not null"`
	Weight     string    `gorm:"not null"`
	Score      int       `gorm:"not null"`
	Active     bool      `gorm:"default:false"`
	Activeat   time.Time `gorm:""`
	Inactiveat time.Time `gorm:""`
	Title      string    `gorm:"default:'__title__'"`
	Exp        int       `gorm:"default:1"`
}

type Screenshot struct {
	gorm.Model
	UID       string `gorm:""`
	UserID    uint   `gorm:"not null"`
	MissionID uint   `gorm:"not null"`
	Picture   string `gorm:"not null"`
	Audit     bool   `gorm:"default:false"`
	Approve   bool   `gorm:"default:false"`
}

type Assignment struct {
	gorm.Model
	UserID    uint `gorm:"not null"`
	MissionID uint `gorm:"not null"`
}

type Emailtoken struct {
	gorm.Model
	UserID   uint      `gorm:"not null"`
	Token    string    `gorm:"unique_index;not null"`
	Expireat time.Time `gorm:"not null"`
	Usedat   time.Time `gorm:""`
}
