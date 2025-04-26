package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nmarniesse/food-advisor/menu"
)

func getWeekMenu(w http.ResponseWriter, r *http.Request) {
	food := []string{"jambon", "tomate", "poulet"}
	query := &menu.Query{food, 30, true, 3}

	ia := &menu.ChatGPT{os.Getenv("CHATGPT_TOKEN")}
	_, err := ia.RunQuery(query)

	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintln(w, "End")
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("APP_ENV") == "dev" {
		log.Println("Listening on http://localhost:8080")
	}

	http.HandleFunc("/week-menu", getWeekMenu)
	http.ListenAndServe(":8080", nil)
}
