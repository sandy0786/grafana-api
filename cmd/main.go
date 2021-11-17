package main

import (
	"log"
	"net/http"

	constants "github.com/sandy0786/grafana-api/constants"
	data "github.com/sandy0786/grafana-api/data"
	handler "github.com/sandy0786/grafana-api/handlers"

	"github.com/gorilla/mux"
)

func init() {
	//set data
	data.SetData()
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc(constants.STATUS_PATH, handler.Status).Methods("GET")
	myRouter.HandleFunc(constants.SUGGEST_PATH, handler.Suggest).Methods("GET")
	myRouter.HandleFunc(constants.TAGNAMES_PATH, handler.TagNames).Methods("GET")
	myRouter.HandleFunc(constants.FILTERS_PATH, handler.Filters).Methods("GET")
	myRouter.HandleFunc(constants.AGGREGATIONS_PATH, handler.Aggregations).Methods("GET")
	myRouter.HandleFunc(constants.QUERY_PATH, handler.Query).Methods("GET")
	myRouter.HandleFunc(constants.DATA_PATH, handler.Data).Methods("GET")
	// myRouter.HandleFunc(constants.DOC_PATH, handler.DocHandler).Methods("GET")
	fs := http.FileServer(http.Dir("F:\\Tutorial\\Backend\\grafana-api\\swaggerui\\"))
	myRouter.PathPrefix(constants.DOC_PATH).Handler(http.StripPrefix(constants.DOC_PATH, fs))
	log.Println("grafana-api started at 8080 port")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
