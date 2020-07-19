package initdb

import (
	"fmt"
	"lil-helper-backend/config"
	"lil-helper-backend/db"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDatabase() {

	config := config.Config.Database

	// payload := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	payload := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
	DB, err := gorm.Open("postgres", payload)
	if err != nil {
		panic("Database initialization failed: " + err.Error())
	} else {
		db.LilHelperDB = DB
		db.LilHelperDB.DB().SetMaxIdleConns(10)
		db.LilHelperDB.DB().SetMaxOpenConns(100)
	}
}
