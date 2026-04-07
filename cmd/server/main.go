package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	
	"github.com/gattini0928/Equilibrium/internal/db"
	userHandlerPkg  "github.com/gattini0928/Equilibrium/internal/handlers/users"
	userRepositoryPkg "github.com/gattini0928/Equilibrium/internal/repositories/users"
	"github.com/gattini0928/Equilibrium/internal/routes"
	userServicePkg "github.com/gattini0928/Equilibrium/internal/services/users"
	configs "github.com/gattini0928/Equilibrium/internal/configs"
)

func main() {
	err := godotenv.Load()
    if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
    }

	conn := db.Connect()

	cfg := configs.LoadDBConfig()
	secret := []byte(cfg.JWTSecret)

	userRepo := userRepositoryPkg.NewUserRepository(conn)
	userService := userServicePkg.NewUserService(userRepo, secret)
	userHandler := userHandlerPkg.NewUserHandler(userService) 

	mux := http.NewServeMux()

	routes.UserRoutes(mux, userHandler, secret)
	
	port := os.Getenv("API_PORT")
	err = http.ListenAndServe(":"+port, mux)
	
	if err != nil {
		log.Fatal(err)
	}
}






