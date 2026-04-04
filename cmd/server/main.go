package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
    }
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to Equilibrium"))
	})
	port := os.Getenv("API_PORT")
	http.ListenAndServe(":"+port, mux)
}






