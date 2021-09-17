package main

import (
	"os"

	"pasarwarga-service-api/database"
	"pasarwarga-service-api/migration"

	"github.com/gin-gonic/gin"
)

func main() {
	command := os.Args[1]
	database.InitDB()

	if command == "migrate" {
		migration.Migrate()
	} else {
		r := gin.Default()

		initRoutes(r)

		r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}
}
