## LAZY DOC

🌩️ Lazy doc is a little tool that help developper with API documentation

Made with golang 

---

### Usage :

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
myos@teste:~$ lazy_doc -generate
```

This command will generate a api_doc.html and a style.css for you 😎