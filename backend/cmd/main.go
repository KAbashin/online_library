package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	// "online_library/backend/internal/handlers"
	"online_library/backend/internal/routes"

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
	r.Use(cors.Default()) // УБРАТЬ В ПРОДЕ !!!!
	/*
		r.Use(cors.New(cors.Config{
			AllowOrigins: []string{
				"http://localhost:5173",
				"http://localhost:5174",
				"http://localhost:5175",
				"http://localhost:5176",
				"http://localhost:5177",
				"http://localhost:5178",
				"http://localhost:5179",
				"http://localhost:5180",
			},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
		}))
	*/

	routes.SetupRoutes(r, db)

	log.Println("Сервер запущен на :8080")
	errServ := http.ListenAndServe(":8080", r)
	if errServ != nil {
		log.Fatal("Не удалось запустить сервер:", errDb)
		return
	}
}
