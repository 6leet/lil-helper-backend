package initdb

import (
	"fmt"
	"lil-helper-backend/db"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "leolee"
	password = "password"
	dbname   = "lil-helper-db"
)

func InitDatabase() {
	payload := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DB, err := gorm.Open("postgres", payload)
	if err != nil {
		panic("Database initialization failed: " + err.Error())
	} else {
		db.LilHelperDB = DB
		db.LilHelperDB.DB().SetMaxIdleConns(10)
		db.LilHelperDB.DB().SetMaxOpenConns(100)
	}
}
