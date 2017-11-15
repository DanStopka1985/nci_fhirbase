package dao

import (
	"net/http"
	"fmt"
	"database/sql"
	ss "../settings"
)

var s = ss.GetSettings()
var fhirPsqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", s.DbFhirHost, s.DbFhirPort, s.DbFhirUser, s.DbFhirPassword, s.DbFhirName)
var fhirDb, errFhirCon = sql.Open("postgres", fhirPsqlInfo)


func setHeaders(w http.ResponseWriter) http.ResponseWriter  {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return w
}

func getRows(query string) (*sql.Rows) {
	var err error
	err = errFhirCon
	P(err)
	P(fhirDb.Ping())
	P(err)
	rows, err := fhirDb.Query(query)
	P(err)
	return rows
}

func P(err error){
	if err != nil {
		panic(err)
	}
}

type ResourceSearch struct {
	ResourceType string `json:"resourceType"`
	QueryString string `json:"queryString"`
}

type ResourceSelect struct {
	ResourceType string `json:"resourceType"`
	Id string `json:"id"`
	VersionId string `json:"versionId"`
}