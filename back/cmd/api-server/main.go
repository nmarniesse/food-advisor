package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nmarniesse/food-advisor/api"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("APP_ENV") == "dev" {
		log.Println("Listening on http://localhost:8080")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/week-menu", api.GetWeekMenu)
	mux.HandleFunc("/refine-week-menu", api.GetWeekMenuRefined)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
