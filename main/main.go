package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	_ "github.com/lib/pq"
	ss "../settings"
	rep "../dao"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/nci/fhir/{resourceType}", rep.GetResourceSearchResult)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(ss.GetSettings().Port), router))
}