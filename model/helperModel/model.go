package helpermodel

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UID      string `gorm:""`
	Username string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
}
