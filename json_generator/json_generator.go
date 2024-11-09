package jsongenerator

import (
	"encoding/json"
	"fmt"
	"lazy_doc/modele"
	"lazy_doc/utilities"
	"os"
)

func write_json_api(filename string, value modele.Lazy_doc) error{
	data, err := json.MarshalIndent(value, "","    ")

	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func Jsongenerator(filename string) {
	fmt.Println("Lazy doc json generator")
	var api_json modele.Lazy_doc

	api_json.Name = utilities.Texte_insertion("Name")
	api_json.Description = utilities.Texte_insertion("Description")
	api_json.Version = utilities.String_insertion("Version")

	api_json.Endpoints = utilities.Endpoint_insertion()


	fmt.Println(api_json)
	write_json_api(filename, api_json)
}
