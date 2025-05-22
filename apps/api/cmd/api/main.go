package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"partiq/internal/proposals"
	"partiq/internal/shared"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"

	"partiq/internal/processes"
)

func main() {
	db, err := shared.ConnectDB()

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("❌ PORT is not set in .env")
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
	proposalRepo := proposals.NewRepository(db)
	proposalSvc := proposals.NewService(proposalRepo)
	proposalHandler := proposals.NewHandler(proposalSvc)

	r := chi.NewRouter()
	origins := strings.Split(os.Getenv("CORS_ORIGINS"), ",")
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	r.Route("/processes", func(r chi.Router) {
		r.Get("/", processHandler.GetAll)
		r.Route("/{processID}", func(r chi.Router) {
			r.Get("/proposals", proposalHandler.GetByProcessID)
		})
	})

	log.Printf("🚀 starting server on :%s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("❌ server failed: %v", err)
	}
}
