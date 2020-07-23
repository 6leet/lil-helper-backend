package inittable

import (
	helpermodel "lil-helper-backend/model/helperModel"

	"github.com/jinzhu/gorm"
)

func MigrateTable(DB *gorm.DB) {
	DB.AutoMigrate(
		helpermodel.User{},
		helpermodel.Mission{},
		helpermodel.Screenshot{},
		helpermodel.Assignment{},
		helpermodel.Emailtoken{},
	)
}
