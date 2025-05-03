package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/nmarniesse/food-advisor/menu"
	"github.com/rs/cors"
)

func getWeekMenu(w http.ResponseWriter, r *http.Request) {
	var err error
	foodInFridgeParam := r.FormValue("food_in_fridge")
	var foodInFridge []string
	if foodInFridgeParam == "" {
		foodInFridge = []string{}
	} else {
		foodInFridge = strings.Split(foodInFridgeParam, ",")
	}

	maxPreparationTime, err := strconv.Atoi(r.FormValue("max_preparation_time"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "max_preparation_time must be an integer")
		return
	}

	persons, err := strconv.Atoi(r.FormValue("persons"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "persons must be an integer")
		return
	}

	useSasonIngredients, err := strconv.ParseBool(r.FormValue("use_season_ingredients"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "use_season_ingredients must be a boolean")
		return
	}

	log.Println("foodInFridge:", foodInFridge)
	log.Println("maxPreparationTime:", maxPreparationTime)
	log.Println("persons:", persons)
	log.Println("useSasonIngredients:", useSasonIngredients)

	query := &menu.Query{
		FoodInFridge:            foodInFridge,
		MaxPreparationTimeInMin: maxPreparationTime,
		UseSeasonIngredient:     useSasonIngredients,
		Persons:                 persons,
	}

	isFake := os.Getenv("FAKE_AI") == "1"
	var menus *menu.Response
	if isFake {
		fake := menu.Fake{}
		menus, _ = fake.RunQuery(query)
	} else {
		ia := &menu.ChatGPT{Token: os.Getenv("CHATGPT_TOKEN")}
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

func getWeekMenuRefined(w http.ResponseWriter, r *http.Request) {
	var err error

	daysToKeepParam := r.FormValue("days_to_keep")
	var daysToKeep []string
	if daysToKeepParam == "" {
		daysToKeep = []string{}
	} else {
		daysToKeep = strings.Split(daysToKeepParam, ",")
	}

	query := &menu.RefineQuery{
		DaysToKeep: daysToKeep,
	}

	isFake := os.Getenv("FAKE_AI") == "1"
	var menus *menu.Response
	if isFake {
		fake := menu.Fake{}
		menus, _ = fake.RunRefineQuery(query)
	} else {
		ia := &menu.ChatGPT{Token: os.Getenv("CHATGPT_TOKEN")}
		menus, err = ia.RunRefineQuery(query)
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
	mux.HandleFunc("/refine-week-menu", getWeekMenuRefined)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
