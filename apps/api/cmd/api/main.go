package main

import (
	"database/sql"
	"log"
	"net/http"
	"partiq/internal/shared"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"

	"partiq/internal/processes"
)

func main() {
	db, err := shared.ConnectDB()

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("failed to close DB: %v\n", err)
		}
	}(db)

	processRepo := processes.NewRepository(db)
	processSvc := processes.NewService(processRepo)
	processHandler := processes.NewHandler(processSvc)

	r := chi.NewRouter()
	r.Get("/processes", processHandler.GetAll)

	log.Println("Server running on :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
