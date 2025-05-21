package main

import (
	"log"
	"partiq/internal/shared"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db, err := shared.ConnectDB()
	if err != nil {
		log.Fatalf("❌ failed to connect to DB: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("❌ failed to close DB: %v", err)
		}
	}()

	var userID int
	err = db.QueryRow(`
		INSERT INTO users (email, name)
		VALUES ($1, $2)
		RETURNING id
	`, "jan.kowalski@example.com", "Jan Kowalski").Scan(&userID)

	if err != nil {
		log.Fatalf("❌ failed to insert user: %v", err)
	}

	var processID int
	err = db.QueryRow(`
		INSERT INTO processes (title, description, start_at, end_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, "Budżet obywatelski 2025", "Zgłoś swoje pomysły na zmiany w mieście!", time.Now(), time.Now().AddDate(0, 1, 0)).Scan(&processID)

	if err != nil {
		log.Fatalf("❌ failed to insert process: %v", err)
	}

	_, err = db.Exec(`
		INSERT INTO proposals (process_id, user_id, title, body)
		VALUES ($1, $2, $3, $4)
	`, processID, userID, "Ścieżka rowerowa przy ul. Zielonej", "Proszę o budowę ścieżki rowerowej przy ul. Zielonej. Poprawi to bezpieczeństwo mieszkańców.")

	if err != nil {
		log.Fatalf("❌ failed to insert proposal: %v", err)
	}

	log.Println("✅ Seed data inserted: 1 user, 1 process, 1 proposal")
}
