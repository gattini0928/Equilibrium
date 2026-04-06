package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"

	"github.com/gattini0928/Equilibrium/internal/db"
)

func main() {
	godotenv.Load()
	conn := db.Connect()
	defer conn.Close()

	drop := `
	DROP TABLE IF EXISTS consultation_remedies CASCADE;
	DROP TABLE IF EXISTS consultation_books CASCADE;
	DROP TABLE IF EXISTS consultations CASCADE;
	DROP TABLE IF EXISTS agendas CASCADE;
	DROP TABLE IF EXISTS remedies CASCADE;
	DROP TABLE IF EXISTS books CASCADE;
	DROP TABLE IF EXISTS patients CASCADE;
	DROP TABLE IF EXISTS therapists CASCADE;
	DROP TABLE IF EXISTS psychiatrists CASCADE;
	DROP TABLE IF EXISTS users CASCADE;
	`

	_, err := conn.Exec(drop)
	if err != nil {
		log.Fatal(err)
	}

	query := `
		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100),
			email VARCHAR(200) UNIQUE,
			password VARCHAR(255),
			age INTEGER,
			cpf VARCHAR(11),
			role VARCHAR(20),
			image TEXT
		);

		CREATE TABLE patients (
			id SERIAL PRIMARY KEY,
			user_id INTEGER UNIQUE REFERENCES users(id) ON DELETE CASCADE,
			therapist_id INTEGER REFERENCES therapists(id),
			psychiatrist_id INTEGER REFERENCES psychiatrists(id),
			current_diagnosis TEXT
		);

		CREATE TABLE therapists (
			id SERIAL PRIMARY KEY,
			user_id INTEGER UNIQUE REFERENCES users(id) ON DELETE CASCADE,
			specialty TEXT,
			description TEXT
		);

		CREATE TABLE psychiatrists (
			id SERIAL PRIMARY KEY,
			user_id INTEGER UNIQUE REFERENCES users(id) ON DELETE CASCADE,
			crm TEXT,
			description TEXT
		);

		CREATE TABLE consultations (
			id SERIAL PRIMARY KEY,
			patient_id INTEGER REFERENCES patients(id) ON DELETE CASCADE,
			professional_id INTEGER,
			professional_type TEXT,
			date TIMESTAMP,
			price REAL,
			annotation TEXT
		);

		CREATE TABLE books (
			id SERIAL PRIMARY KEY,
			author TEXT,
			title TEXT
		);

		CREATE TABLE remedies (
			id SERIAL PRIMARY KEY,
			name TEXT,
			dosage TEXT,
			quantity INTEGER
		);

		CREATE TABLE consultation_books (
			consultation_id INTEGER REFERENCES consultations(id) ON DELETE CASCADE,
			book_id INTEGER REFERENCES books(id) ON DELETE CASCADE
		);

		CREATE TABLE consultation_remedies (
			consultation_id INTEGER REFERENCES consultations(id) ON DELETE CASCADE,
			remedy_id INTEGER REFERENCES remedies(id) ON DELETE CASCADE
		);

		CREATE TABLE agendas (
			id SERIAL PRIMARY KEY,
			professional_id INTEGER,
			day INTEGER,
			month INTEGER,
			hour TEXT,
			reserved BOOLEAN
		);
	`

	_, err = conn.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Banco criado corretamente 🚀")
}