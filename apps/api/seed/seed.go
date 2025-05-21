package seed

import (
	"github.com/joho/godotenv"
	"log"
	"partiq/internal/shared"
	"time"
)

func main() {
	_ = godotenv.Load()
	db, err := shared.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("failed to close DB: %v", err)
		}
	}()

	_, err = db.Exec(`
		INSERT INTO processes (title, description, start_at, end_at)
		VALUES ($1, $2, $3, $4)
	`, "Budżet obywatelski 2025", "Zgłoś swoje pomysły na zmiany w mieście!", time.Now(), time.Now().AddDate(0, 1, 0))

	if err != nil {
		log.Fatalf("failed to insert seed data: %v", err)
	}

	log.Println("✅ Seed data inserted successfully.")
}
