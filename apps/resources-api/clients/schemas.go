package clients

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

var Schemas map[string]string

func InitSchemas() {
	resp, err := http.Get("http://localhost:3000/schemas")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var schemas []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&schemas); err != nil {
		panic(err)
	}

	schemasMap := make(map[string]string)
	for _, schema := range schemas {
		schemaId := schema["$id"].(string)
		resourceName := strings.Split(schemaId, "/")[len(strings.Split(schemaId, "/"))-1]
		resourceName = strings.TrimSuffix(resourceName, ".schema")
		schemaJson, _ := json.Marshal(schema)
		schemasMap[resourceName] = string(schemaJson)
	}

	Schemas = schemasMap
	log.Println(Schemas)
}
