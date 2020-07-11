package inittable

import (
	"lil-helper-backend/model/helperModel"

	"github.com/jinzhu/gorm"
)

func MigrateTable(DB *gorm.DB) { //
	DB.AutoMigrate(
		&helperModel.User{},
	)
}
