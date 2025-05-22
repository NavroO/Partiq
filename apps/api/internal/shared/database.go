package shared

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func ConnectDB() (*sql.DB, error) {
	_ = godotenv.Load()
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Error().Msg("❌ DATABASE_URL is not set")
		return nil, ErrEnvNotSet("DATABASE_URL")
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Error().Err(err).Msg("❌ failed to open DB connection")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Error().Err(err).Msg("❌ failed to ping DB")
		return nil, err
	}

	log.Info().Msg("✅ Connected to database")
	return db, nil
}

type ErrEnvNotSet string

func (e ErrEnvNotSet) Error() string {
	return "environment variable not set: " + string(e)
}
