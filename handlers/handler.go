package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	constants "github.com/sandy0786/grafana-api/constants"
	data "github.com/sandy0786/grafana-api/data"
	entities "github.com/sandy0786/grafana-api/entities"
	service "github.com/sandy0786/grafana-api/service"
	validations "github.com/sandy0786/grafana-api/validations"
)

// swagger:route GET / users APIStatus
// Get Status of the api
//
// responses:
//  200: APIStatus
func Status(w http.ResponseWriter, r *http.Request) {
	log.Println("inside status ")
	status := entities.Status{
		Status:  "ok",
		Message: "Datasource is working",
	}
	w.Header().Set("Content-Type", "application/json") // set content type
	w.Header().Set("Access-Control-Allow-Origin", "*") // set content type
	w.WriteHeader(http.StatusOK)                       // return status code
	json.NewEncoder(w).Encode(status)
}

// swagger:route GET /suggest users listDistinctValues
// Get distinct column values
//
// responses:
//  200: SuggestList
//  500: CommonError
func Suggest(w http.ResponseWriter, r *http.Request) {
	log.Println("inside Suggest ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(constants.SUGGEST_CONSTANTS)
}

// swagger:route GET /tagnames users listColumnNames
// Get column names list
//
// responses:
//  500: CommonError
//  200: GetColumnsList
func TagNames(w http.ResponseWriter, r *http.Request) {
	log.Println("inside TagNames ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(constants.SUGGEST_CONSTANTS)
}

// swagger:route GET /filters users listFilterColumns
// Get filter column list
//
// responses:
//  500: CommonError
//  200: FiltersList
func Filters(w http.ResponseWriter, r *http.Request) {
	log.Println("inside TagNames ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(constants.FILTER_CONSTANTS)
}

// swagger:route GET /aggregations users listAggregateFuncions
// Get Aggregate functions list
//
// responses:
//  500: CommonError
// 	200: AggregationList
func Aggregations(w http.ResponseWriter, r *http.Request) {
	log.Println("inside Aggregations ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(constants.AGGREGATE_FUNCTIONS)
}

// swagger:operation GET /query users listDataByFilters
// Get graphable data by applying filters
//
// ---
// parameters:
// - name: start
//   in: query
//   description: start time in epoch milliseconds
//   type: integer
//   required: true
// - name: end
//   in: query
//   description: end time in epoch milliseconds
//   type: integer
//   required: true
// - name: columns
//   in: query
//   description: comma separated column names
//   type: string
//   required: true
// - name: filter
//   in: query
//   description: filter in {columnName=value} format. default filter is contains
//   type: string
//   required: false
// responses:
//   "200":
//     "$ref": "#/responses/DataList"
//   "500":
//     "$ref": "#/responses/CommonError"
//   "400":
//     "$ref": "#/responses/ValidationError"
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

// swagger:route GET /data users listCompleteData
// Get complete data
//
// responses:
//  200: DataList
func Data(w http.ResponseWriter, r *http.Request) {
	log.Println("inside Data ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Data)
}
