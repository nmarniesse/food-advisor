package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nmarniesse/food-advisor/menu"
	"github.com/rs/cors"
)

func getWeekMenu(w http.ResponseWriter, r *http.Request) {
	food := []string{"jambon", "tomate", "poulet", "pates", "oeufs"}
	query := &menu.Query{food, 30, true, 3}

	isFake := os.Getenv("FAKE_AI") == "1"
	var menus *menu.Response
	var err error
	if isFake {
		fake := menu.Fake{}
		menus, _ = fake.RunQuery(query)
	} else {
		ia := &menu.ChatGPT{os.Getenv("CHATGPT_TOKEN")}
		menus, err = ia.RunQuery(query)
		if err != nil {
			log.Panicln(err)
		}
	}

	log.Println(menus)
	response, err := json.Marshal(menus)
	if err != nil {
		log.Panicln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response))
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("APP_ENV") == "dev" {
		log.Println("Listening on http://localhost:8080")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/week-menu", getWeekMenu)

	handler := cors.Default().Handler(mux)
    http.ListenAndServe(":8080", handler)
}
