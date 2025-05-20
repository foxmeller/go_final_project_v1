package main

import (
	//"fmt"

	"github.com/foxmeller/go_final_project_v1/db"
	"github.com/foxmeller/go_final_project_v1/httpServer"
)

func main() {
	db.DbExistance()
	httpServer.StartWebServer()

}
