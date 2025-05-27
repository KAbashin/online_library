package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	// "online_library/backend/internal/handlers"
	"online_library/backend/internal/routes"
	// "online_library/backend/migrations"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Запустили проект")
	db, errDb := sql.Open("postgres", "postgres://librarian:pass@localhost:5432/mydb?sslmode=disable")
	if errDb != nil {
		log.Fatalf("Не удалось подключиться к БД:  %v", errDb)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// возможно это надо выносить в отдельный скрипт для периодического опроса
	if err := db.Ping(); err != nil {
		log.Fatalf("Не удалось подключиться к БД (ping не прошёл): %v", err)
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
