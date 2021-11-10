package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	constants "github.com/sandy0786/grafana-api/constants"
	entities "github.com/sandy0786/grafana-api/entities"
	service "github.com/sandy0786/grafana-api/service"
	validations "github.com/sandy0786/grafana-api/validations"
)

func Status(w http.ResponseWriter, r *http.Request) {
	log.Println("inside status ")
	status := entities.Status{
		Status:  "ok",
		Message: "Datasource is working",
	}
	w.Header().Set("Content-Type", "application/json") // set content type
	w.WriteHeader(http.StatusOK)                       // return status code
	json.NewEncoder(w).Encode(status)
}

func Suggest(w http.ResponseWriter, r *http.Request) {
	log.Println("inside Suggest ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(constants.SUGGEST_CONSTANTS)
}

func Aggregations(w http.ResponseWriter, r *http.Request) {
	log.Println("inside Aggregations ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(constants.AGGREGATE_FUNCTIONS)
}

func Query(w http.ResponseWriter, r *http.Request) {
	log.Println("inside Query ")

	queryParams := r.URL.Query()
	isValid, err := validations.ValidateRequest(queryParams)
	if isValid {
		response := service.GetData(queryParams)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if len(response) > 0 {
			json.NewEncoder(w).Encode(response)
		} else {
			json.NewEncoder(w).Encode([]map[string]interface{}{})
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

}
