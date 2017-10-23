package reports

import (
	"net/http"
	"encoding/json"
	"database/sql"
	"fmt"
)


type ResourceSearch struct {
	ResourceType string `json:"resourceType"`
	QueryString string `json:"queryString"`
}

func GetResourceSearchResult(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	resourceSearch := &ResourceSearch{ResourceType: "CodeSystem", QueryString:`name=foo`}
	b, err := json.Marshal(resourceSearch)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	var in []byte


	//rows := getRows(`SET plv8.start_proc = 'plv8_init'; select fhir_search::jsonb val from fhir_search('{"resourceType": "CodeSystem", "queryString": ""}');`)
	rows := getRows(`SET plv8.start_proc = 'plv8_init'; select fhir_search::jsonb val from fhir_search('` + string(b) + `');`)
	for rows.Next() {
		var val sql.NullString
		rows.Scan(&val)
 		in = []byte(val.String)
	}
	var raw map[string]interface{}
	json.Unmarshal(in, &raw)
	out, _ := json.Marshal(raw)
	w.Write(out)
}
