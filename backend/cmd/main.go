package cmd

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"online_library/backend/internal/handlers"
	"online_library/backend/internal/routes"
	"online_library/backend/migrations"
)

func main() {

	r := gin.Default()
	routes.SetupRoutes(r)

	log.Fatal(r.Run(":8080"))
}
