package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/nmarniesse/food-advisor/internal/di"
	"github.com/nmarniesse/food-advisor/internal/model"
)

func GetWeekMenu(w http.ResponseWriter, r *http.Request) {
	di := di.NewDI()
	defer di.Shutdown()

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

	query := &model.Query{
		FoodInFridge:            foodInFridge,
		MaxPreparationTimeInMin: maxPreparationTime,
		UseSeasonIngredient:     useSasonIngredients,
		Persons:                 persons,
	}

	menus, err := di.Ia.RunQuery(query)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(menus)
	response, err := json.Marshal(menus)
	if err != nil {
		log.Panicln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response))
}

func GetWeekMenuRefined(w http.ResponseWriter, r *http.Request) {
	di := di.NewDI()
	defer di.Shutdown()

	var err error

	uuid, err := uuid.Parse(r.FormValue("uuid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "uuid must be provided and valid")
		return
	}

	daysToKeepParam := r.FormValue("days_to_keep")
	var daysToKeep []string
	if daysToKeepParam == "" {
		daysToKeep = []string{}
	} else {
		daysToKeep = strings.Split(daysToKeepParam, ",")
	}

	query := &model.RefineQuery{Uuid: uuid, DaysToKeep: daysToKeep}
	menus, err := di.Ia.RunRefineQuery(query)
	if err != nil {
		log.Panicln(err)
	}

	response, err := json.Marshal(menus)
	if err != nil {
		log.Panicln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response))
}
