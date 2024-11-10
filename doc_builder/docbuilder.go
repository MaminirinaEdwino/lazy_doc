package docbuilder

import (
	"lazy_doc/modele"
	"os"
	"strings"
)

func Doc_builder(lazy_doc modele.Lazy_doc, html_doc bool, md_doc bool) {
	if html_doc {
		template := Html_template()
		template = strings.Replace(template, "api_name", lazy_doc.Name, 2)
		template = strings.Replace(template, "api_description", lazy_doc.Description, 1)
		template = strings.Replace(template, "api_version", lazy_doc.Version, 1)
		template = strings.Replace(template, "api_endpoints", EndpointList(lazy_doc.Endpoints), 1)
		os.WriteFile("api_doc.html", []byte(template), 0644)
		Generate_css()
	}
	if md_doc {
		MdTemplate := MdTemplate()
		MdTemplate = strings.Replace(MdTemplate, "api_name", lazy_doc.Name, 1)
		MdTemplate = strings.Replace(MdTemplate, "api_description", lazy_doc.Description, 1)
		MdTemplate = strings.Replace(MdTemplate, "api_version", lazy_doc.Version, 1)
		MdTemplate = strings.Replace(MdTemplate, "api_endpoint", MdEndpointTemplate(lazy_doc.Endpoints), 1)

		os.WriteFile("api_doc.md", []byte(MdTemplate), 0644)
	}

}
