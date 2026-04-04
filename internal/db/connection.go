package db

import (
	"database/sql"
	"log"
	"github.com/gattini0928/Equilibrium/internal/configs"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	cfg := configs.LoadDBConfig()
	connStr := cfg.ConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}