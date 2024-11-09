package utilities

import (
	"bufio"
	"fmt"
	"lazy_doc/modele"
	"os"
)
func isrootvalid(root string) bool{
	isvalid := false
	if root[0] == '/' {
		isvalid = true
	}
	return isvalid
}

func root_choice() string {
	root := String_insertion("root")
	if !isrootvalid(root) {
		fmt.Println("Add a valid root : /example/com")
		root = root_choice()
	}
	return root
}

func isvalidmethod(method string) bool {
	var method_list [5]string
	method_list[0] = "GET"
	method_list[1] = "POST"
	method_list[2] = "PUT"
	method_list[3] = "PATCH"
	method_list[4] = "DELETE"
	isvalid := false

	for _, value := range method_list {
		if !isvalid {
			if value == method {
				isvalid = true
			}
		}
	}
	return isvalid
}

func method_choice() string {
	method := String_insertion("method : ")
	if !isvalidmethod(method) {
		fmt.Println("Add a valid method : GET POST PUT PATCH DELETE")
		method = method_choice()
	}
	return method
}

func asking_new_end_point() bool {
	var choice string
	choice = String_insertion("Add a new endpoint ? (n/y)")
	if choice == "n" {
		return false
	} else if choice == "y" {
		return true
	} else {
		asking_new_end_point()
	}
	return false
}

func Endpoint_insertion() []modele.Endpoint {
	var endpoints []modele.Endpoint //endpoint list
	var endpoint modele.Endpoint
	fmt.Println("Add endpoints")
	for asking_new_end_point() == true {
		endpoint.Tag = String_insertion("Enter the endpoint tag ")
		endpoint.Root = root_choice()
		endpoint.Method = method_choice()
		if endpoint.Method == "GET" {
			endpoint.Body = ""
		} else {
			endpoint.Body =  Texte_insertion("request body ")
		}
		endpoint.Response = Texte_insertion("response body ")
		endpoint.Status_code = String_insertion("status code ")
		endpoint.Description = Texte_insertion("description ")
		endpoint.Summary = Texte_insertion("summary ")

		endpoints = append(endpoints, endpoint)
	}

	return endpoints
}

func String_insertion(label string) string {
	var variable string
	fmt.Printf("%v : ", label)
	fmt.Scanln(&variable)
	return variable
}

func Texte_insertion(label string) string {
	fmt.Printf("%v : ", label)
	// fmt.Scanln(&variable)
	scanner := bufio.NewScanner(os.Stdin)

	var lines string

	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines += line +"\n"
	}
	return lines
}