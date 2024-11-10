package docbuilder

import (
	"lazy_doc/modele"
	"os"
	"strings"
)

func Html_template() string {
	template := "<!DOCTYPE html><html lang=\"en\"><head>\n<meta charset=\"UTF-8\">\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link rel=\"stylesheet\" href=\"style.css\">\n<title> api_name </title>\n</head>\n<body><header><h3 class=\"lazy_doc\">Lazy doc</h3><h1> api_name</h1><h3> Description api_description </h3><h3> Version api_version</h3> </header> \napi_endpoints </body></html>"
	return template
}

func Endpoint_template() string {
	return "<div class=\"endpoint\">\n<h3>api_tag</h3>\n<div class=\"root border_api_method\">\n<span class=\"api_method method\">api_method</span><span>api_root </span><span class=\"summary\"> api_summary </span></div>\n<div class=\"description\">\nDescription : api_description</div><br><div class=\"request_body\">\nRequest Body <br> <div class=\"code\">\n<code>api_body</code></div>\n</div>\n<div class=\"response_body\"><div class=\"status_code\"><br>Status Code : api_status_code</div><br><div class=\"response\">Response body <br> <div class=\"code\"><code>api_response</code></div></div></div></div>"
}

func Generate_css() {
	template_css := "* {padding: 0;margin: 0;font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;}header{border: 2px solid rgb(0, 108, 231);padding: 10px;border-radius: 5px;}.lazy_doc{font-size: 2.5rem;color: rgb(82, 80, 223);}.endpoint{border: 2px solid rgb(2, 61, 255);border-radius: 5px;padding: 10px;margin: 20px;}.root{border: 1px solid black;display: flex;gap: 10px;align-items: center;padding: 2px;border-radius: 5px;}.method{background-color: red;display: block ;padding: 5px;min-width: 100px;text-align: center;font-weight: 600;color: white;border-radius: 5px;}.summary{font-size: 12px;}.code{background-color: rgb(2, 1, 1);border-radius: 5px;color: white;padding: 10px;}.GET{background-color: rgb(0, 153, 255);}.POST{background-color: rgb(33, 228, 33);}.PUT{background-color: orangered;}.PATCH{background-color: orange;}.DELETE{background-color: rgb(233, 4, 4);}.border_GET{border: 1px solid rgb(0, 153, 255);}.border_POST{border: 1px solid rgb(33, 228, 33);}.border_PUT{border: 1px solid orangered;}.border_PATCH{border: 1px solid orange;}.border_DELETE{    border: 1px solid rgb(233, 4, 4);}"

	os.WriteFile("style.css", []byte(template_css), 0644)
}

func EndpointList(endpoints []modele.Endpoint) string {
	var template_endpoints string
	for _, endpoint := range endpoints {
		template_endpoint := Endpoint_template()
		endpoint.Body = strings.ReplaceAll(endpoint.Body, "\n", "\n<br>")
		endpoint.Response = strings.ReplaceAll(endpoint.Response, "\n", "\n<br>")
		template_endpoint = strings.Replace(template_endpoint, "api_tag", endpoint.Tag, 1)
		template_endpoint = strings.Replace(template_endpoint, "api_method", endpoint.Method, 3)
		template_endpoint = strings.Replace(template_endpoint, "api_root", endpoint.Root, 1)
		template_endpoint = strings.Replace(template_endpoint, "api_summary", endpoint.Summary, 1)
		template_endpoint = strings.Replace(template_endpoint, "api_description", endpoint.Description, 1)
		template_endpoint = strings.Replace(template_endpoint, "api_body", endpoint.Body, 1)

		template_endpoint = strings.Replace(template_endpoint, "api_status_code", endpoint.Status_code, 1)
		template_endpoint = strings.Replace(template_endpoint, "api_response", endpoint.Response, 1)
		template_endpoints += template_endpoint
	}
	return template_endpoints
}

func MdTemplate() string {
	Template := "# api_name api_version \n"+
	"### api_description \n"+
	"## Endpoints : \n"+
	"api_endpoint"
	

	return Template
}

func MdEndpointTemplate(endpoints []modele.Endpoint)string{
	template := "> **api_tag** : \n"+
	"api_method api_root api_summary \n"+
	"Description : api_description \n"+
	"Status code api_status \n"+
	"* Request Body \n"+
	"``` json \n"+
	"api_body \n"+
	"``` \n"+
	"* Response Body \n"+
	"``` json \n"+
	"api_response \n"+
	"``` \n "
	var EndpointList string
	for _, endpoint := range endpoints{
		endp:= template
		endp = strings.ReplaceAll(endp, "api_tag", endpoint.Tag)
		endp = strings.ReplaceAll(endp, "api_method", endpoint.Method)
		endp = strings.ReplaceAll(endp, "api_root", endpoint.Root)
		endp = strings.ReplaceAll(endp, "api_summary", endpoint.Summary)
		endp = strings.ReplaceAll(endp, "api_description", endpoint.Description)
		endp = strings.ReplaceAll(endp, "api_status", endpoint.Status_code)
		
		endp = strings.ReplaceAll(endp, "api_body", endpoint.Body)
		endp = strings.ReplaceAll(endp, "api_response", endpoint.Response)
		EndpointList+= endp
	}


	return EndpointList
}