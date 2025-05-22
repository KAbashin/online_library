package cmd

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"online_library/backend/internal/handlers"
	"online_library/backend/internal/routes"
	"online_library/backend/migrations"

	_ "github.com/lib/pq"
)

func main() {
	db, errDb := sql.Open("postgres", "postgres://user:pass@localhost:5432/mydb?sslmode=disable")
	if errDb != nil {
		log.Fatal("Не удалось подключиться к БД:", errDb)
	}

	r := gin.Default()
	routes.SetupRoutes(r, db)

	log.Println("Сервер запущен на :8080")
	errServ := http.ListenAndServe(":8080", r)
	if errServ != nil {
		log.Fatal("Не удалось запустить сервер:", errDb)
		return
	}
}
