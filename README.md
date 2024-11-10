## LAZY DOC

🌩️ Lazy doc is a little tool that help developper with API documentation

Made with golang 

---

### Usage :

#### Method 1 : 

1. **Create a json file named api.json in your project directory**

2. **Add some data like this example**

```json
{
  "Name": "The name of your API\n",
  "Version": "The Version",
  "Description": "The description of your API\n",
  "Endpoints": [
    {
      "Tag": "The tag name for the Endpoint",
      "Root": "/GET",
      "Method": "GET",
      "Body": "",
      "Response": "{\n\t\"endpoint\":\"teste\"\n}\n",
      "Status_code": "200",
      "Description": "Get all the doc\n",
      "Summary": "Summary\n"
    }
  ]
}
```

#### Explication :

- **Name** : This key Contains the name of your API
- **Version** : is for the version
- **Descritpion** : you can describe your API in this section

- **Endpoints** : in this part, you can add all the endpoints that you need, you just need to follow the example

```json
{
  "Tag": "The tag name for the Endpoint",
  "Root": "/GET",
  "Method": "GET",
  "Body": "",
  "Response": "{\n\t\"endpoint\":\"teste\"\n}\n",
  "Status_code": "200",
  "Description": "Get all the doc\n",
  "Summary": "Summary\n"
}
```

3. **Generate the html doc file**

* use the command : 
```console 
lazy_doc -generate -html -md #or 
lazy_doc -generate -html #or
lazy_doc -generate -md
```

This command will generate a api_doc.html and a style.css for you 😎

#### Method 2 : 
1. use the command below to generate an api.json file :
```console
lazy doc -create_doc
```
2. complete all thing you want
3. use command below to generate the doc : 
```console
lazy_doc -generate -md -html #or
lazy_doc -generate -html #or
lazy_doc -generate -md
```