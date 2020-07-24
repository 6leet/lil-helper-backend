package main

import (
	"flag"
	"lil-helper-backend/db"
	initdb "lil-helper-backend/init/initDB"
	initrouter "lil-helper-backend/init/initRouter"
	inittable "lil-helper-backend/init/initTable"
	"strconv"
)

// @title lil-helper swagger API
// @version 0.0.1
// @discription API for lil-helper
// @name Authorization
// @BasePath /backend
func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "ip port (int)")
	flag.Parse()

	// fmt.Println(config.Config.Mission.Weights)
	// fmt.Println(config.VTool.GetInt("mission.maxlevel"))
	// for i := 0; i <= config.Config.Mission.Maxlevel; i++ {
	// 	config.Config.Mission.Weights[i]++
	// }
	// config.VTool.Set("mission.weights", config.Config.Mission.Weights)
	// fmt.Println(config.Config.Mission.Maxlevel)
	// fmt.Println(config.VTool.GetInt("mission.maxlevel"))
	// config.VTool.WriteConfig()
	// fmt.Println(config.Config.Mission.Weights)
	// fmt.Println(config.VTool.GetInt("mission.maxlevel"))
	initdb.InitDatabase()

	inittable.MigrateTable(db.LilHelperDB)

	router := initrouter.InitRouter()
	router.Run(":" + strconv.Itoa(port))
}
