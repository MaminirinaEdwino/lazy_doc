package filereader

import (
	"encoding/json"
	"fmt"
	"lazy_doc/modele"
	"os"
)

func Read_api_file() modele.Lazy_doc {
	data, err1:= os.ReadFile("api.json")
	if err1 != nil {
		fmt.Println("Create the api file first !")
	}
	var lazy_doc modele.Lazy_doc
	err := json.Unmarshal(data, &lazy_doc)
	if err != nil {
		panic(err)
	}
	return lazy_doc
}

