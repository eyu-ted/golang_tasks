// main.go
package main

import (
	"task_manager/data"
	"task_manager/router"


)

func main() {
	data.ConnectDB()
	data.InitUserCollection()
	r := router.SetupRouter()
	r.Run(":8080")
}
