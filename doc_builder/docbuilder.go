package docbuilder

import (
	"lazy_doc/modele"
	"os"
	"strings"
)

func Doc_builder(lazy_doc modele.Lazy_doc)  {
	template:=Html_template()
	template = strings.Replace(template, "api_name", lazy_doc.Name, 2)
	template = strings.Replace(template, "api_description", lazy_doc.Description, 1)
	template = strings.Replace(template, "api_version", lazy_doc.Version, 1)
	template = strings.Replace(template, "api_endpoints", EndpointList(lazy_doc.Endpoints), 1)
	os.WriteFile("api_doc.html", []byte(template), 0644)
	Generate_css()
}