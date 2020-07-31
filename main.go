package main

import (
	"flag"
	"lil-helper-backend/db"
	"lil-helper-backend/goroutine"
	initdb "lil-helper-backend/init/initDB"
	initrouter "lil-helper-backend/init/initRouter"
	inittable "lil-helper-backend/init/initTable"
	helpermodel "lil-helper-backend/model/helperModel"
	"strconv"
)

// @title lil-helper swagger API
// @version 0.0.1
// @discription API for lil-helper
// @name Authorization
// @BasePath /backend
func _main() {
	var port int
	flag.IntVar(&port, "port", 8080, "ip port (int)")
	flag.Parse()

	initdb.InitDatabase()

	inittable.MigrateTable(db.LilHelperDB)

	router := initrouter.InitRouter()
	router.Run(":" + strconv.Itoa(port))

	goroutine.Wg.Done()
}

func main() {
	goroutine.Wg.Add(2)
	go _main()
	go helpermodel.AutoReorganizeMission("12:12")
	goroutine.Wg.Wait()
}
