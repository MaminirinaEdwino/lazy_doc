package main

import (
	"flag"
	"fmt"
	docbuilder "lazy_doc/doc_builder"
	filereader "lazy_doc/file_reader"
	jsongenerator "lazy_doc/json_generator"
	// jsongenerator "lazy_doc/json_generator"
)

func main() {
	fmt.Println("Lazy doc")
	// jsongenerator.Jsongenerator("teste.json")
	create := flag.Bool("create_doc", false, "create the api.json file")
	generate := flag.Bool("generate", false, "generate the web documentation")
	help := flag.Bool("help", false, "help")

	flag.Parse()
	
	switch {
	case *create:
		jsongenerator.Jsongenerator("api.json")
	case *generate:
		docbuilder.Html_template()
		docbuilder.Doc_builder(filereader.Read_api_file())
	
	case *help:
		fmt.Println("help")
	default:
		fmt.Println("option -help")
	}

}
