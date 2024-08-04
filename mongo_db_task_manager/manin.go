// main.go
package main

import (
	"task_manager/data"
	"task_manager/router"
	// "log"
	// "github.com/joho/godotenv"

)

func main() {
	data.ConnectDB()
	r := router.SetupRouter()
	r.Run(":8080")
}
