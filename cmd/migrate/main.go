package main

import (
	"log"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"github.com/gattini0928/Equilibrium/internal/db"
)

func main(){
	godotenv.Load()
	conn := db.Connect()
	defer conn.Close()

	dropQuery := `
		DROP TABLE IF EXISTS users CASCADE;
		DROP TABLE IF EXISTS patients CASCADE;
		DROP TABLE IF EXISTS therapists CASCADE;
		DROP TABLE IF EXISTS psychiatrists CASCADE;
	`

	_, err := conn.Exec(dropQuery)
	if err != nil {
		log.Fatal(err)
	}

	query := `
		CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(200) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			age INTEGER NOT NULL,
			cpf VARCHAR(11) NOT NULL,
			role VARCHAR(20) CHECK (role IN ('patient', 'therapist', 'psychiatrist')) NOT NULL,
			image TEXT
		);

		CREATE TABLE IF NOT EXISTS patients(
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE
		);

		CREATE TABLE IF NOT EXISTS therapists(
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
			specialty VARCHAR(100) NOT NULL,
			description TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS psychiatrists(
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
			crm VARCHAR(20) NOT NULL
		);
	`
	_, err = conn.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tabelas criadas com sucesso")
}
