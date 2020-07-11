package main

import (
	"flag"
	"lil-helper-backend/db"
	initdb "lil-helper-backend/init/initDB"
	initrouter "lil-helper-backend/init/initRouter"
	inittable "lil-helper-backend/init/initTable"
	"strconv"
)

// @title swagger Example API (lil-helper)
// @version 0.0.1
// @discription API for lil-helper
// @name Authorization
// @BasePath
func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "ip port (int)")
	flag.Parse()

	initdb.InitDatabase()

	inittable.MigrateTable(db.LilHelperDB)

	router := initrouter.InitRouter()
	router.Run(":" + strconv.Itoa(port))
}
