package main

import (
	"github.com/aryansh13/go_restapi_gin/database"
	"github.com/aryansh13/go_restapi_gin/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
